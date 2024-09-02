package main

import (
	"github.com/DuganChandler/pokedexgo/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	player        Player
	nextUrl       *string
	prevUrl       *string
}
