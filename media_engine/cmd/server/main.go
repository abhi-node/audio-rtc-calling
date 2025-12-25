package main

import (
	"context"
	"fmt"
	"net/http"
	"rtc_media_engine/internal/config"
	"rtc_media_engine/internal/database"
	"rtc_media_engine/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := config.NewConfig()
	repo := repository.NewRepository(config.CONN_STRING)

	err := database.Migrate(context.Background(), repo.Pool)
	if err != nil {
		fmt.Println(err)
		return
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.PORT)
}
