package main

import (
	"sync"

	"github.com/DuganChandler/pokedexgo/internal/pokeapi"
)

type Player struct {
	pokedex map[string]pokeapi.RespPokemonInfo
	mu      *sync.Mutex
}
