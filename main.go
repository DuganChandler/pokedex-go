package main

import (
	"sync"
	"time"

	"github.com/DuganChandler/pokedexgo/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		player: Player{
			pokedex: make(map[string]pokeapi.RespPokemonInfo),
			mu:      &sync.Mutex{},
		},
		pokeapiClient: pokeClient,
	}

	startREPL(cfg)
}
