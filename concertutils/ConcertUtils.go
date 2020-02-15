package concertutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func IsBandAlreadyInSlice(concertSlice []Concert, bandName string) bool {

	for _, concert := range concertSlice {
		if strings.EqualFold(concert.Artist, bandName) {
			return true
		}
	}

	return false
}

func GetSongKickAPIKey() string {
	bs, err := ioutil.ReadFile("keys/SongKickApiKey.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(bs))
}

func GetLastFMAPIKey() string {
	bs, err := ioutil.ReadFile("keys/LastFMApiKey.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(bs))
}

func RemoveDuplicateEvents(concerts []Concert, concertsToAdd []Concert) []Concert {
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

func RemoveAccents(word string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, word)

	return s
}
