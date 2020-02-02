package main

import (
	"net/http"
	"sort"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	songKickAPIKey := getSongKickAPIKey()
	lastFMAPIKey := getLastFMAPIKey()

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/concerts/:area/:user", func(c *gin.Context) {
		artists := getMostListenedArtists(c.Param("user"), lastFMAPIKey, c.Query("limit"))
		concerts := readConcertsInArea(c.Param("area"), lastFMAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcerts/:user", func(c *gin.Context) {
		artists := getMostListenedArtists(c.Param("user"), lastFMAPIKey, c.Query("limit"))
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540", "56332", "28796"}

		concerts := getConcertsForUser(skAreaSlice, songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcertsAllUsers/:user1/:user2", func(c *gin.Context) {
		artistsUser1 := getMostListenedArtists(c.Param("user1"), lastFMAPIKey, c.Query("limit"))
		artistsUser2 := getMostListenedArtists(c.Param("user2"), lastFMAPIKey, c.Query("limit"))
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540", "56332", "28796"}

		allConcerts := getConcertsForUser(skAreaSlice, songKickAPIKey, artistsUser1)
		concertsSecondUser := getConcertsForUser(skAreaSlice, songKickAPIKey, artistsUser2)

		allConcerts = removeDuplicateEvents(allConcerts, concertsSecondUser)

		sort.Slice(allConcerts, func(i, j int) bool {
			return allConcerts[i].Date.Before(allConcerts[j].Date)
		})

		c.JSON(http.StatusOK, allConcerts)
	})

	r.Run(":8282")
}
