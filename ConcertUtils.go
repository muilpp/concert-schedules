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

func getSongKickApiKey() string {
	bs, err := ioutil.ReadFile("SongKickApiKey.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return string(bs)
}

func getLastFMApiKey() string {
	bs, err := ioutil.ReadFile("LastFMApiKey.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return string(bs)
}
