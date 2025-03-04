package handlers

import (
	"net/http"
	"time"
	"url-shortener-api/models"
	"url-shortener-api/utils"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context, urlShortener *models.URLShortener) {
	var req struct {
		OriginalURL string `json:"original_url"`
	}

	// Parse JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	urlShortener.Mutex.Lock()
	defer urlShortener.Mutex.Unlock()

	// Check if the URL is already shortened
	for short, original := range urlShortener.Store {
		if original == req.OriginalURL {
			c.JSON(http.StatusOK, gin.H{"short_url": short})
			return
		}
	}

	// Generate a new short URL
	shortURL := utils.GenerateShortURL()
	urlShortener.Store[shortURL] = req.OriginalURL

	// Update ClickStore and LogData
	domain := utils.GetDomain(req.OriginalURL)
	urlShortener.ClickStore[domain]++

	urlShortener.LogData = append(urlShortener.LogData, models.Log{
		CreatedAt:  time.Now().Format(time.RFC3339),
		ShortenUrl: shortURL,
	})

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}
