package concertutils

import (
	"strings"
	"testing"
	"time"
)

func TestIsBandAlreadyInSliceWhenFound(t *testing.T) {
	var concertSlice []Concert

	concert := CreateConcert("Event", time.Now(), "artist", "venue", "city")
	concertSlice = append(concertSlice, *concert)
	repeatedConcert := CreateConcert("Event", time.Now(), "artist", "venue", "city")

	if !IsBandAlreadyInSlice(concertSlice, repeatedConcert.Artist) {
		t.Errorf("Expected band %v to be in slice, but it's not", concert.Artist)
	}
}

func TestIsBandAlreadyInSliceWhenNotFound(t *testing.T) {
	var concertSlice []Concert

	concert := CreateConcert("Event", time.Now(), "artist", "venue", "city")
	concertSlice = append(concertSlice, *concert)
	repeatedConcert := CreateConcert("Event", time.Now(), "different artist", "venue", "city")

	if IsBandAlreadyInSlice(concertSlice, repeatedConcert.Artist) {
		t.Errorf("Expected band %v not to be in slice, but it is", concert.Artist)
	}
}

func TestRemoveDuplicateEvents(t *testing.T) {
	var concertSlice []Concert
	concertSlice = append(concertSlice, *CreateConcert("Event", time.Now(), "artist1", "venue", "city"))
	concertSlice = append(concertSlice, *CreateConcert("Event", time.Now(), "artist2", "venue", "city"))
	concertSlice = append(concertSlice, *CreateConcert("Event", time.Now(), "artist3", "venue", "city"))

	var duplicateConcertSlice []Concert
	duplicateConcertSlice = append(concertSlice, *CreateConcert("Event", time.Now(), "artist2", "venue", "city"))
	duplicateConcertSlice = append(concertSlice, *CreateConcert("Event", time.Now(), "artist3", "venue", "city"))
	duplicateConcertSlice = append(concertSlice, *CreateConcert("Event", time.Now(), "artist4", "venue", "city"))

	cleanConcertSlice := RemoveDuplicateEvents(concertSlice, duplicateConcertSlice)

	totalElementsExpected := 4
	if len(cleanConcertSlice) != totalElementsExpected {
		t.Errorf("Duplicates are not removed properly from concert slice, expected %v elements, but %v found", totalElementsExpected, len(cleanConcertSlice))
	}
}

func TestRemoveAccents(t *testing.T) {
	word := RemoveAccents("Bonć diáàñ!")

	if !strings.Contains(word, "Bonc diaan") {
		t.Errorf("No special characters expected, but %v found", word)
	}
}
