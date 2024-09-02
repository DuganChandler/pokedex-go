package main

import (
	"github.com/DuganChandler/pokedexgo/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextUrl       *string
	prevUrl       *string
}
