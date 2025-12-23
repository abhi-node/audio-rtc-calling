package main

import (
	"net/http"
	"rtc_media_engine/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := config.NewConfig()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.PORT)
}
