package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortURL(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func HandleShortURLRedirect(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
