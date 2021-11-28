package concertutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/marc/concerts/artistutils"
)

func GetConcertsInOneArea(area string, apiKey string, artistSlice []artistutils.Artist) []Concert {
	c := make(chan []byte)

	count := 0
	for i := 0; i < 20; i++ {
		go requestConcertsInArea(area, apiKey, strconv.Itoa(i), c)
		count++
	}

	var concertArray []Concert
	for i := 0; i < 20; i++ {
		var response = <-c

		if len(response) == 0 {
			//Got until the last request
			break
		}

		var jsonResponse JSONResponse
		json.Unmarshal([]byte(response), &jsonResponse)

		for _, event := range jsonResponse.ResultsPage.Results.Event {

			var performanceSlice []Performance
			for _, performance := range event.Performance {
				performance.Artist = RemoveAccents(performance.Artist)
				performanceSlice = append(performanceSlice, performance)
			}

			for _, artist := range artistSlice {
				addConcert := false
				var bandName string

				for _, performance := range performanceSlice {
					if strings.EqualFold(artist.Name, performance.Artist) && !IsBandAlreadyInSlice(concertArray, performance.Artist) {
						addConcert = true
					}

					if addConcert {
						bandName = performance.Artist
						break
					}
				}

				if addConcert {
					t, err := time.Parse("2006-01-02", event.Start.Date)

					if err != nil {
						log.Println(err)
						break
					}

					concert := CreateConcert(event.Name, t, bandName, event.Venue.Name, event.Location.City)
					concertArray = append(concertArray, *concert)
				}
			}
		}
	}

	return concertArray
}

func GetConcertsInMultipleAreas(skAreaSlice []string, songKickAPIKey string, artists []artistutils.Artist) []Concert {
	concertChannel := make(chan []Concert)
	for _, skArea := range skAreaSlice {
		go getConcertsInAreaWithChannel(skArea, songKickAPIKey, artists, concertChannel)
	}

	var concerts []Concert
	for i := 0; i < len(skAreaSlice); i++ {
		newConcerts := <-concertChannel

		for _, concert := range newConcerts {
			if !IsBandAlreadyInSlice(concerts, concert.Artist) {
				concerts = append(concerts, concert)
			}
		}
	}

	return concerts
}

func requestConcertsInArea(area string, apiKey string, page string, c chan []byte) {
	client := http.Client{}

	Url, err := url.Parse("https://api.songkick.com")
	Url.Path += "/api/3.0/metro_areas/" + area + "/calendar.json"
	parameters := url.Values{}
	parameters.Add("apikey", apiKey)
	parameters.Add("page", page)
	Url.RawQuery = parameters.Encode()

	req, err := http.NewRequest("GET", Url.String(), nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Got until request number %v", page)
		var emptyByteSlice []byte
		c <- emptyByteSlice
	}

	response, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading response")
	}

	c <- response
}

func getConcertsInAreaWithChannel(area string, apiKey string, artistSlice []artistutils.Artist, c chan []Concert) {
	c <- GetConcertsInOneArea(area, apiKey, artistSlice)
}
