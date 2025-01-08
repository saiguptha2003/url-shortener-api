package main

import (
	"url-shortener-api/handlers"
	"url-shortener-api/models"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	urlShortener := models.NewURLShortener()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
			"endpoints": []string{
				"/shorten - POST request to shorten a URL",
				"/:shortURL - GET request to redirect to the original URL",
				"/top-domains - GET request to fetch top domains",
			},
			"note": "Use /shorten to shorten URLs. For more details, check the other endpoints.",
		})
	})
	r.POST("/shorten", func(c *gin.Context) {
		handlers.ShortenURL(c, urlShortener)
	})
	r.GET("/:shortURL", func(c *gin.Context) {
		handlers.RedirectToOriginalURL(c, urlShortener)
	})

	r.GET("/top-domains", func(c *gin.Context) {
		handlers.GetTopDomains(c, urlShortener)
	})

	return r
}


func main() {
	r := SetupRouter()
	r.Run(":5000")
}
