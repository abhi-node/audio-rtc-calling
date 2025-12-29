package main

import (
	"context"
	"rtc_media_engine/internal/config"
	"rtc_media_engine/internal/handler"
	"rtc_media_engine/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := config.NewConfig()
	handler := handler.NewHandler(*config, context.Background())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "X-Request-Id"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)

	protected := r.Group("/api")
	protected.Use(middleware.AuthGuard(config))
	{
		protected.POST("/friends/create", handler.SendFriendRequest)
		protected.PUT("/friends/accept", handler.AcceptFriendRequest)
		protected.GET("/friends/incoming", handler.GetIncomingRequests)
		protected.GET("/friends/outgoing", handler.GetOutgoingRequests)
		protected.GET("/friends", handler.GetFriends)
	}

	r.Run(config.PORT)
}
