package handler

import (
	"net/http"
	"rtc_media_engine/internal/dto"
	"rtc_media_engine/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var registerBody dto.RegisterRequest
	err := c.BindJSON(&registerBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body"})
	}

	user := models.User{
		Email: registerBody.Email,
	}
	token, err := h.s.RegisterUser(&user, registerBody.Password, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) LoginHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var loginBody dto.LoginRequest
	err := c.BindJSON(&loginBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body"})
	}

	user := models.User{
		Email: loginBody.Email,
	}
	token, err := h.s.LoginUser(&user, loginBody.Password, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
