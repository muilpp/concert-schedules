package main

import (
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/joho/godotenv"
	"github.com/marc/concerts/userutils"

	"github.com/marc/concerts/concertutils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load("credentials.env")

	if err != nil {
		log.Fatalf("Error loading credentials.env file")
	}

	songKickAPIKey := os.Getenv("SONG_KICK_API_KEY")

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/concerts/:area/:user/:lastFMAPIKey", func(c *gin.Context) {
		user := userutils.CreateNewUser(c.Param("user"))
		artists := user.GetMostListenedArtists(c.Param("lastFMAPIKey"), c.Query("limit"))
		concerts := concertutils.GetConcertsInOneArea(c.Param("area"), songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcerts/:user/:lastFMAPIKey/:areas", func(c *gin.Context) {
		user := userutils.CreateNewUser(c.Param("user"))
		artists := user.GetMostListenedArtists(c.Param("lastFMAPIKey"), c.Query("limit"))
		skAreaSlice := strings.Split(c.Param("areas"), ",")
		concerts := concertutils.GetConcertsInMultipleAreas(skAreaSlice, songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcertsAllUsers/:user1/:user2/:lastFMAPIKey/:areas", func(c *gin.Context) {
		user1 := userutils.CreateNewUser(c.Param("user1"))
		user2 := userutils.CreateNewUser(c.Param("user2"))
		artistsUser1 := user1.GetMostListenedArtists(c.Param("lastFMAPIKey"), c.Query("limit"))
		artistsUser2 := user2.GetMostListenedArtists(c.Param("lastFMAPIKey"), c.Query("limit"))
		skAreaSlice := strings.Split(c.Param("areas"), ",")

		concertsFirstUser := concertutils.GetConcertsInMultipleAreas(skAreaSlice, songKickAPIKey, artistsUser1)
		concertsSecondUser := concertutils.GetConcertsInMultipleAreas(skAreaSlice, songKickAPIKey, artistsUser2)
		allConcerts := concertutils.RemoveDuplicateEvents(concertsFirstUser, concertsSecondUser)

		sort.Slice(allConcerts, func(i, j int) bool {
			return allConcerts[i].Date.Before(allConcerts[j].Date)
		})

		c.JSON(http.StatusOK, allConcerts)
	})

	r.Run(":8282")
}
