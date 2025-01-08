package main

import (
	"url-shortener-api/handlers"
	"url-shortener-api/models"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	urlShortener := models.NewURLShortener()
	r := gin.Default()

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
	r.Run(":8080")
}
