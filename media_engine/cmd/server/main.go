package main

import (
	"context"
	"rtc_media_engine/internal/config"
	"rtc_media_engine/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := config.NewConfig()
	handler := handler.NewHandler(*config, context.Background())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)

	r.Run(config.PORT)
}
