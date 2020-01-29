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
		artists := getMostListenedArtists(c.Param("user"), lastFMAPIKey)

		concerts := readConcertsInArea(c.Param("area"), songKickAPIKey, artists)

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.GET("/allConcerts/:user", func(c *gin.Context) {
		artists := getMostListenedArtists(c.Param("user"), lastFMAPIKey)
		//TODO Read this from file or database
		skAreaSlice := []string{"28714", "28480", "28539", "28604", "28540"}

		concertChannel := make(chan []Concert)
		for _, skArea := range skAreaSlice {
			go readConcertsInAreaByUser(skArea, songKickAPIKey, artists, concertChannel)
		}

		var concerts []Concert
		for i := 0; i < len(skAreaSlice); i++ {
			newConcerts := <-concertChannel

			for _, concert := range newConcerts {
				if !isBandAlreadyInSlice(concerts, concert.Artist) {
					concerts = append(concerts, concert)
				}
			}
		}

		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})

		c.JSON(http.StatusOK, concerts)
	})

	r.Run(":8282")
}
