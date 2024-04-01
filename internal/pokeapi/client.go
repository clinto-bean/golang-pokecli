package pokeapi

import (
	"net/http"
	"time"

	"github.com/clinto-bean/golang-pokecli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

func CreateClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}