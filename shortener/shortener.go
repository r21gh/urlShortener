// This file implements the shortening function of the project
// I can create a module based on mathematics to use base62 
// which contains alphabet (26 small char, 26 capital char, and 10 digits)
// but the sake of simplicity we try teris-io package
package shortener

import (
	"fmt"
	"github.com/teris-io/shortid"
	"urlShortener/storage"
)

// Shortener is the interface that wraps UrlShortening operations
type Shortener interface {
	Shorten(url string) (string, error)
	Get(key string) (string, error)
}

// Store shortens and persists Urls
// Note: Don't really love this name
type Store struct {
	storage storage.Storage
}

// TODO: Consider making a shortId generator specific to
// each urlShortener struct
var generateID = func() (string, error) {
	return shortid.Generate()
}

// Shorten creates a shortened URL Key and persists the key to the underlying store
// Returns the shortened url key if successful
func (u *Store) Shorten(url string) (string, error) {
	id, err := generateID()
	if err != nil {
		return "", fmt.Errorf("generating id for %s: %v", url, err)
	}

	exists, err := u.storage.Exists(id)
	if err != nil {
		return "", fmt.Errorf("checking existance of key %s for url %s: %v", id, url, err)
	} else if exists {
		return "", fmt.Errorf("key %s for url %s already exists in storage", id, url)
	}

	if err := u.storage.Set(id, url); err != nil {
		return "", fmt.Errorf("storing key %s for url %s: %v", id, url, err)
	}

	return id, nil
}

// Get fetches the URL by it's shortened key if it exists
// Returns an empty string if not found
func (u Store) Get(key string) (string, error) {
	return u.storage.Get(key)
}

// New creates and returns a new Store
func New(storage storage.Storage) *Store {
	return &Store{storage}
}
