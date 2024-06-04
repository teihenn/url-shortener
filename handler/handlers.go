package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teihenn/url-shortener/shortener"
	"github.com/teihenn/url-shortener/store"
)

// Request model definition
type URLCreationRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var request URLCreationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := shortener.GenerateShortLink(request.LongURL, request.UserID)
	store.SaveURLMapping(shortURL, request.LongURL, request.UserID)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortURLRedirect(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
