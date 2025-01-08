package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateShortURL() string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)
	rand.Seed(time.Now().UnixNano())

	for i := range shortURL {
		shortURL[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(shortURL)
}

func GetDomain(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 2 {
		return parts[2]
	}
	return "unknown"
}
