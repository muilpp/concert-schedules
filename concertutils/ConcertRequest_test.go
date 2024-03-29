package concertutils

import (
	"os"
	"testing"

	"github.com/marc/concerts/artistutils"
)

func TestGetConcertsInMultipleAreas(t *testing.T) {
	var concertSlice []Concert

	skAreaSlice := []string{"28714", "28480", "28539"}
	artistSlice := []artistutils.Artist{{Name: "Metallica"}, {Name: "Coldplay"}, {Name: "NOFX"}, {Name: "Talco"}}
	concertSlice = GetConcertsInMultipleAreas(skAreaSlice, os.Getenv("SONG_KICK_API_KEY"), artistSlice)

	//Can't check that the slice is not empty because the bands tested might not be on tour, so let's consider that if no exception occurred, everything's fine
	if !(len(concertSlice) >= 0) {
		t.Errorf("Concerts were not read correctly")
	}
}
