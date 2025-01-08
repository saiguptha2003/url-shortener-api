package handlers

import (
	"net/http"
	"sort"
	"url-shortener-api/models"

	"github.com/gin-gonic/gin"
)

func GetTopDomains(c *gin.Context, urlShortener *models.URLShortener) {
	urlShortener.Mutex.RLock()
	defer urlShortener.Mutex.RUnlock()

	type domainCount struct {
		Domain string
		Count  int
	}

	// Collect domain counts
	var domainCounts []domainCount
	for domain, count := range urlShortener.ClickStore {
		domainCounts = append(domainCounts, domainCount{Domain: domain, Count: count})
	}

	// Sort by count in descending order
	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].Count > domainCounts[j].Count
	})

	// Get top 3 domains
	topDomains := domainCounts
	if len(domainCounts) > 3 {
		topDomains = domainCounts[:3]
	}

	c.JSON(http.StatusOK, gin.H{"top_domains": topDomains})
}
