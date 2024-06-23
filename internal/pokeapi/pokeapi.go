package pokeapi

import (
	"net/http"
	"time"
)

const baseURI = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
