package handler

import (
	"net/http"
	"rtc_media_engine/internal/dto"
	"rtc_media_engine/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendFriendRequest(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody dto.SendFriendRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Improper JSON format"})
		return
	}

	val, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	userId, ok := val.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	user := models.User{
		UserID: userId,
	}
	friend := models.User{
		Email: requestBody.Email,
	}

	err = h.s.SendFriendRequest(&user, &friend, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully sent request"})
}

func (h *Handler) AcceptFriendRequest(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody dto.AcceptFriendRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Improper JSON format"})
		return
	}

	val, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	userId, ok := val.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	user := models.User{
		UserID: userId,
	}
	friend := models.User{
		UserID: requestBody.FriendId,
	}

	err = h.s.AcceptFriendRequest(&user, &friend, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully sent request"})
}

func (h *Handler) GetIncomingRequests(c *gin.Context) {
	ctx := c.Request.Context()
	val, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	userId, ok := val.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	user := models.User{
		UserID: userId,
	}

	users, err := h.s.GetIncomingRequests(&user, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": users})
}

func (h *Handler) GetOutgoingRequests(c *gin.Context) {
	ctx := c.Request.Context()
	val, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	userId, ok := val.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	user := models.User{
		UserID: userId,
	}

	users, err := h.s.GetOutgoingRequests(&user, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": users})
}

func (h *Handler) GetFriends(c *gin.Context) {
	ctx := c.Request.Context()
	val, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	userId, ok := val.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Improper Token Format"})
		return
	}

	user := models.User{
		UserID: userId,
	}

	users, err := h.s.GetFriends(&user, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": users})
}
