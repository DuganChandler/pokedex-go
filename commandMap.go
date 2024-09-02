package main

import (
	"fmt"
)

// Display next 20 areas
func commandMap(cfg *config, str string) error {
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextUrl)
	if err != nil {
		return err
	}
	cfg.nextUrl = locationsResponse.NextURL
	cfg.prevUrl = locationsResponse.PrevURL

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

// Display previous areas
func commandMapb(cfg *config, str string) error {
	if cfg.prevUrl == nil {
		return fmt.Errorf("you are on the first page of areas, please use 'map' to pull up new ones")
	}

	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevUrl)
	if err != nil {
		return err
	}
	cfg.nextUrl = locationsResponse.NextURL
	cfg.prevUrl = locationsResponse.PrevURL

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
