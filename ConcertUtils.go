package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isBandAlreadyInSlice(concertSlice []Concert, bandName string) bool {

	for _, concert := range concertSlice {
		if strings.EqualFold(concert.Artist, bandName) {
			return true
		}
	}

	return false
}

func getSongKickAPIKey() string {
	bs, err := ioutil.ReadFile("SongKickApiKey.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(bs))
}

func getLastFMAPIKey() string {
	bs, err := ioutil.ReadFile("LastFMApiKey.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(bs))
}

func getConcertsForUser(skAreaSlice []string, songKickAPIKey string, artists []Artist) []Concert {
	concertChannel := make(chan []Concert)
	for _, skArea := range skAreaSlice {
		go readConcertsInAreaByUser(skArea, songKickAPIKey, artists, concertChannel)
	}

	var concerts []Concert
	for i := 0; i < len(skAreaSlice); i++ {
		newConcerts := <-concertChannel

		for _, concert := range newConcerts {
			if !isBandAlreadyInSlice(concerts, concert.Artist) {
				concerts = append(concerts, concert)
			}
		}
	}

	return concerts
}

func removeDuplicateEvents(concerts []Concert, concertsToAdd []Concert) []Concert {
	for _, concertUser2 := range concertsToAdd {
		isConcertInList := false
		for _, concertUser1 := range concerts {
			if concertUser1.Artist == concertUser2.Artist {
				isConcertInList = true
			}
		}

		if !isConcertInList {
			concerts = append(concerts, concertUser2)
		}
	}

	return concerts
}
