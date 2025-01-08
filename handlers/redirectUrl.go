package handlers

import (
	"net/http"
	"url-shortener-api/models"

	"github.com/gin-gonic/gin"
)

func RedirectToOriginalURL(c *gin.Context, urlShortener *models.URLShortener) {
	shortURL := c.Param("shortURL")

	urlShortener.Mutex.RLock()
	defer urlShortener.Mutex.RUnlock()

	// Redirect if the short URL exists
	if originalURL, exists := urlShortener.Store[shortURL]; exists {
		c.Redirect(http.StatusMovedPermanently, originalURL)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
}
