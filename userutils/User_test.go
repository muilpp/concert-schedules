package userutils

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/marc/concerts/concertutils"
)

func TestGetMostListenedArtists(t *testing.T) {

	bs, err := ioutil.ReadFile("../keys/UserTest.txt")

	if err != nil {
		t.Errorf("Error reading user from file: %v", err)
	}

	user := CreateNewUser(strings.TrimSpace(string(bs)))
	artistSlice := user.GetMostListenedArtists(concertutils.GetLastFMAPIKey(), "150")

	if !(len(artistSlice) > 0) {
		t.Errorf("No bands found for user %v", user.userName)
	}
}
