package main

import (
	"time"
	"github.com/kryptonn36/pokedexcli/internal/pokeapi"
	"github.com/kryptonn36/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(10*time.Second)
	pokeClient := pokeapi.NewClient(5 * time.Second,cache)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}

