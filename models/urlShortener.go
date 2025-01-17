package models

import (
	"sync"

)

type URLShortener struct {
	Store      map[string]string
	ClickStore map[string]int
	Mutex      sync.RWMutex
	LogData    []Log   

}
type Log struct {
	CreatedAt string
	ShortenUrl string

}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		Store:      make(map[string]string),
		ClickStore: make(map[string]int),
		LogData:    []Log{},
	}
}

