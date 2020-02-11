package main

import (
	"concert-schedules/artistutils"
	"concert-schedules/concertutils"
	"net/http"
	"sort"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	songKickAPIKey := concertutils.GetSongKickAPIKey()
	lastFMAPIKey := concertutils.GetLastFMAPIKey()

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/concerts/:area/:user", func(c *gin.Context) {
		artists := artistutils.GetMostListenedArtists(c.Param("user"), lastFMAPIKey, c.Query("limit"))
		concerts := concertutils.ReadConcertsInArea(c.Param("area"), songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcerts/:user", func(c *gin.Context) {
		artists := artistutils.GetMostListenedArtists(c.Param("user"), lastFMAPIKey, c.Query("limit"))
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540", "56332", "28796"}

		concerts := concertutils.GetConcertsForUser(skAreaSlice, songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcertsAllUsers/:user1/:user2", func(c *gin.Context) {
		artistsUser1 := artistutils.GetMostListenedArtists(c.Param("user1"), lastFMAPIKey, c.Query("limit"))
		artistsUser2 := artistutils.GetMostListenedArtists(c.Param("user2"), lastFMAPIKey, c.Query("limit"))
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540", "56332", "28796"}

		allConcerts := concertutils.GetConcertsForUser(skAreaSlice, songKickAPIKey, artistsUser1)
		concertsSecondUser := concertutils.GetConcertsForUser(skAreaSlice, songKickAPIKey, artistsUser2)

		allConcerts = concertutils.RemoveDuplicateEvents(allConcerts, concertsSecondUser)

		sort.Slice(allConcerts, func(i, j int) bool {
			return allConcerts[i].Date.Before(allConcerts[j].Date)
		})

		c.JSON(http.StatusOK, allConcerts)
	})

	r.Run(":8282")
}
