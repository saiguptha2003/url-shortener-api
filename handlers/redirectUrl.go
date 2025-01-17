package handlers

import (
	"net/http"
	"time"
	"url-shortener-api/models"
	"github.com/gin-gonic/gin"
)

func RedirectToOriginalURL(c *gin.Context, urlShortener *models.URLShortener) {
	shortURL := c.Param("shortURL")

	urlShortener.Mutex.Lock()
	defer urlShortener.Mutex.Unlock()

	// Redirect if the short URL exists
	if originalURL, exists := urlShortener.Store[shortURL]; exists {
		// Update HitOrNot in the log entry
		for i := range urlShortener.LogData {
			if urlShortener.LogData[i].ShortenUrl == shortURL {
				urlShortener.LogData[i].CreatedAt = time.Now().Format(time.RFC3339) // Reset timestamp
				break
			}
		}

		c.Redirect(http.StatusMovedPermanently, originalURL)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
}
