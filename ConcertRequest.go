package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func requestConcertsInArea(area string, apiKey string, page string, c chan []byte) {
	client := http.Client{}

	Url, err := url.Parse("https://api.songkick.com")
	Url.Path += "/api/3.0/metro_areas/" + area + "/calendar.json"
	parameters := url.Values{}
	log.Printf("Api key -> %v", apiKey)
	parameters.Add("apikey", apiKey)
	parameters.Add("page", page)
	Url.RawQuery = "apikey=" + apiKey + "&page=" + page
	// parameters.Encode()

	log.Println(Url.String())

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

func readConcertsInArea(area string, apiKey string, artistSlice []Artist) []Concert {
	return getConcerts(area, apiKey, artistSlice)
}

func readConcertsInAreaByUser(area string, apiKey string, artistSlice []Artist, c chan []Concert) {
	c <- getConcerts(area, apiKey, artistSlice)
}

func getConcerts(area string, apiKey string, artistSlice []Artist) []Concert {
	c := make(chan []byte)

	count := 0
	for i := 0; i < 20; i++ {
		go requestConcertsInArea(area, apiKey, strconv.Itoa(i), c)
		count++
	}

	var concertArray []Concert
	for i := 0; i < 20; i++ {
		var response = <-c
		log.Printf("Mida response %v", len(response))
		if len(response) == 0 {
			//Got until the last request
			break
		}

		var jsonResponse JSONResponse
		json.Unmarshal([]byte(response), &jsonResponse)

		for _, event := range jsonResponse.ResultsPage.Results.Event {
			var concert Concert

			for _, artist := range artistSlice {
				addConcert := false

				for _, performance := range event.Performance {
					if !isBandAlreadyInSlice(concertArray, performance.Artist) && strings.EqualFold(artist.Name, performance.Artist) {
						addConcert = true
					}

					if addConcert {
						concert.Artist = performance.Artist
						break
					}
				}

				if addConcert {
					concert.City = event.Location.City

					t, err := time.Parse("2006-01-02", event.Start.Date)

					if err != nil {
						fmt.Println(err)
					} else {
						concert.Date = t
						// concert.Date = t.In(time.Local).Format("January 02, 2006 (MST)")
					}

					concert.Event = event.Name
					concert.Venue = event.Venue.Name
					concertArray = append(concertArray, concert)
				}
			}
		}
	}

	return concertArray
}
