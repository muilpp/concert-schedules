package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	songKickAPIKey := getSongKickApiKey()
	lastFMAPIKey := getLastFMApiKey()

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/concerts/:area/:user", func(c *gin.Context) {
		artists := getMostListenedArtists(c.Param("user"), lastFMAPIKey)
		concerts := readConcertsInArea(c.Param("area"), songKickAPIKey, artists)

		c.JSON(http.StatusOK, concerts)
	})

	r.Run(":8282")
}
