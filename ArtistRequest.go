package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getMostListenedArtists(user string, apiKey string, limit string) []Artist {
	client := http.Client{}

	req, err := http.NewRequest("GET", "http://ws.audioscrobbler.com/2.0/?method=user.gettopartists&user="+user+"&api_key="+apiKey+"&format=json&limit="+limit, nil)

	if err != nil {
		log.Fatal("Error creating request for most listened artists")
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error getting most listened artists")
	}

	response, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading response")
	}

	var artistDTO ArtistDTO
	json.Unmarshal([]byte(response), &artistDTO)

	return artistDTO.Topartists.Artist
}
