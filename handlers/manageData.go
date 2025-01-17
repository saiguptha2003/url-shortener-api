package handlers

import (
	"log"
	"time"
	"url-shortener-api/models"
	"url-shortener-api/utils"
)

//**********************************************************************************************************************
// As my assumption in the Interview
// taking a bool variable may not work so that i changed the logic now this will work with the createdat time
// Only createdat time will be checked on the removelogdata and for each redirect the time will be updated in the createdAt time
//**********************************************************************************************************************


func RemoveLogData(urlShortener *models.URLShortener) {
	ticker := time.NewTicker(1 * time.Minute)  
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			urlShortener.Mutex.Lock()

			now := time.Now()
			newLogData := []models.Log{}
			for _, logEntry := range urlShortener.LogData {
				createdAt, _ := time.Parse(time.RFC3339, logEntry.CreatedAt)
				if now.Sub(createdAt) < 24*time.Hour {  
					newLogData = append(newLogData, logEntry)
				} else {
					if _, exists := urlShortener.Store[logEntry.ShortenUrl]; exists {
						domain := utils.GetDomain(urlShortener.Store[logEntry.ShortenUrl])

						if count, found := urlShortener.ClickStore[domain]; found {
							if count > 1 {
								urlShortener.ClickStore[domain] = count - 1
								log.Println("Decremented click count for", domain)
							} else {
								delete(urlShortener.ClickStore, domain)
								log.Println("Removed", domain, "from clickstore")
							}
						}
					}

					delete(urlShortener.Store, logEntry.ShortenUrl)
				}
			}

			urlShortener.LogData = newLogData

			urlShortener.Mutex.Unlock()
			log.Println("Cleanup completed.")
		}
	}
}