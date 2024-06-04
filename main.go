package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/teihenn/url-shortener/handler"
	"github.com/teihenn/url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortURLRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
