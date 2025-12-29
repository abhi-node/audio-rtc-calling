package middleware

import (
	"net/http"
	"rtc_media_engine/internal/auth"
	"rtc_media_engine/internal/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthGuard(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := auth.ValidateToken(tokenString, []byte(config.JWT_SECRET))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		c.Set("userId", userId)
		c.Next()

	}
}
