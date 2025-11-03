package pokeapi

import (
	"net/http"
	"time"
	"github.com/kryptonn36/pokedexcli/internal/pokecache"
)

type Client struct{
	httpclient http.Client
	cache *pokecache.Cache
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client{
	return Client{
		httpclient: http.Client{Timeout: timeout,},
		cache: cache,
	}
}