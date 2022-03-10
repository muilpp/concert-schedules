package userutils

import (
	"os"
	"testing"
)

func TestGetMostListenedArtists(t *testing.T) {

	user := CreateNewUser("test")
	artistSlice := user.GetMostListenedArtists(os.Getenv("SONG_KICK_API_KEY"), "150")

	if !(len(artistSlice) > 0) {
		t.Errorf("No bands found for user %v", user.userName)
	}
}
