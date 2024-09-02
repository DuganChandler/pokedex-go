package pokeapi

import (
	"net/http"
	"time"

	"github.com/DuganChandler/pokedexgo/internal/caching"
)

// Client:
type Client struct {
	cache      caching.Cache
	httpClient http.Client
}

// New Client:
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: caching.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
