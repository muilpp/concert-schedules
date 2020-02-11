package artistutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetMostListenedArtists(user string, apiKey string, limit string) []Artist {
	client := http.Client{}

	URL, err := url.Parse("http://ws.audioscrobbler.com")
	URL.Path += "/2.0/"
	parameters := url.Values{}
	parameters.Add("method", "user.gettopartists")
	parameters.Add("user", user)
	parameters.Add("api_key", apiKey)
	parameters.Add("format", "json")
	parameters.Add("limit", limit)
	URL.RawQuery = parameters.Encode()

	req, err := http.NewRequest("GET", URL.String(), nil)

	if err != nil {
		log.Println(err)
		log.Fatal("Error creating request for most listened artists")
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
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
