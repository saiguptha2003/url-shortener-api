package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	router := SetupRouter()

	// Test data
	requestBody := map[string]string{"url": "https://github.com/saiguptha2003?tab=repositories"}
	jsonBody, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "shortURL")
}

func TestRedirectURL(t *testing.T) {
	router := SetupRouter()

	// Shorten the URL
	requestBody := map[string]string{"url": "https://github.com/saiguptha2003?tab=repositories"}
	jsonBody, _ := json.Marshal(requestBody)

	shortenReq, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	shortenReq.Header.Set("Content-Type", "application/json")
	shortenRes := httptest.NewRecorder()

	router.ServeHTTP(shortenRes, shortenReq)

	var shortenResponse map[string]string
	_ = json.Unmarshal(shortenRes.Body.Bytes(), &shortenResponse)
	shortURL := shortenResponse["shortURL"]

	// Test redirect
	redirectReq, _ := http.NewRequest("GET", shortURL, nil)
	redirectRes := httptest.NewRecorder()

	router.ServeHTTP(redirectRes, redirectReq)

	assert.Equal(t, http.StatusTemporaryRedirect, redirectRes.Code)
	assert.Equal(t, "https://github.com/saiguptha2003?tab=repositories", redirectRes.Header().Get("Location"))
}

func TestTopDomains(t *testing.T) {
	router := SetupRouter()

	// Add multiple URLs
	urls := []string{
		"https://example.com",
		"https://example.com",
		"https://youtube.com",
		"https://youtube.com",
		"https://youtube.com",
		"https://wikipedia.org",
	}

	for _, url := range urls {
		requestBody := map[string]string{"url": url}
		jsonBody, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
	}

	// Test top domains
	req, _ := http.NewRequest("GET", "/top-domains", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]int
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(response))
	assert.Equal(t, 3, response["youtube.com"])
	assert.Equal(t, 2, response["example.com"])
	assert.Equal(t, 1, response["wikipedia.org"])
}
