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
		concerts := readConcertsInArea(c.Param("area"), songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcerts/:user", func(c *gin.Context) {
		artists := getMostListenedArtists(c.Param("user"), lastFMAPIKey, c.Query("limit"))
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540"}

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
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540"}

		concerts := getConcertsForUser(skAreaSlice, songKickAPIKey, artistsUser1)
		concerts = append(concerts, getConcertsForUser(skAreaSlice, songKickAPIKey, artistsUser2)...)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.Run(":8282")
}
