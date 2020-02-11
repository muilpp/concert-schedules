package concertutils

import (
	"concert-schedules/artistutils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func RequestConcertsInArea(area string, apiKey string, page string, c chan []byte) {
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

func ReadConcertsInArea(area string, apiKey string, artistSlice []artistutils.Artist) []Concert {
	return GetConcerts(area, apiKey, artistSlice)
}

func ReadConcertsInAreaByUser(area string, apiKey string, artistSlice []artistutils.Artist, c chan []Concert) {
	c <- GetConcerts(area, apiKey, artistSlice)
}

func GetConcerts(area string, apiKey string, artistSlice []artistutils.Artist) []Concert {
	c := make(chan []byte)

	count := 0
	for i := 0; i < 20; i++ {
		go RequestConcertsInArea(area, apiKey, strconv.Itoa(i), c)
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
			for _, artist := range artistSlice {
				addConcert := false
				var bandName string

				for _, performance := range event.Performance {
					if !IsBandAlreadyInSlice(concertArray, performance.Artist) && strings.EqualFold(artist.Name, performance.Artist) {
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