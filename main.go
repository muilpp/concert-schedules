package main

import (
	"net/http"
	"sort"

	"github.com/marc/concerts/userutils"

	"github.com/marc/concerts/concertutils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	songKickAPIKey := concertutils.GetSongKickAPIKey()
	lastFMAPIKey := concertutils.GetLastFMAPIKey()

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/concerts/:area/:user", func(c *gin.Context) {
		user := userutils.CreateNewUser(c.Param("user"))
		artists := user.GetMostListenedArtists(lastFMAPIKey, c.Query("limit"))
		concerts := concertutils.GetConcertsInOneArea(c.Param("area"), songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcerts/:user", func(c *gin.Context) {
		user := userutils.CreateNewUser(c.Param("user"))
		artists := user.GetMostListenedArtists(lastFMAPIKey, c.Query("limit"))
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540", "56332", "28796"}

		concerts := concertutils.GetConcertsInMultipleAreas(skAreaSlice, songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcertsAllUsers/:user1/:user2", func(c *gin.Context) {
		user1 := userutils.CreateNewUser(c.Param("user1"))
		user2 := userutils.CreateNewUser(c.Param("user2"))
		artistsUser1 := user1.GetMostListenedArtists(lastFMAPIKey, c.Query("limit"))
		artistsUser2 := user2.GetMostListenedArtists(lastFMAPIKey, c.Query("limit"))
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540", "56332", "28796"}

		allConcerts := concertutils.GetConcertsInMultipleAreas(skAreaSlice, songKickAPIKey, artistsUser1)
		concertsSecondUser := concertutils.GetConcertsInMultipleAreas(skAreaSlice, songKickAPIKey, artistsUser2)

		allConcerts = concertutils.RemoveDuplicateEvents(allConcerts, concertsSecondUser)

		sort.Slice(allConcerts, func(i, j int) bool {
			return allConcerts[i].Date.Before(allConcerts[j].Date)
		})

		c.JSON(http.StatusOK, allConcerts)
	})

	r.Run(":8282")
}
