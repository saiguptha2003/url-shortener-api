package models

import (
	"sync"
)

type URLShortener struct {
	Store      map[string]string
	ClickStore map[string]int
	Mutex      sync.RWMutex
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		Store:      make(map[string]string),
		ClickStore: make(map[string]int),
	}
}
