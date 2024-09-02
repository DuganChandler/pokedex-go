package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type cliCommand struct {
	name        string
	description string
	callBack    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "brings up help message",
			callBack:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the Pokedex",
			callBack:    commandExit,
		},
		"clear": {
			name:        "clear",
			description: "clears the screen",
			callBack:    commandClear,
		},
		"map": {
			name:        "map",
			description: "shows the next 20 areas to traverse to",
			callBack:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "shows the previous 20 areas to traverse to",
			callBack:    commandMapb,
		},
		"explore": {
			name:        "explore <area-name>",
			description: "shows the pokemon encounters of the given area",
			callBack:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "attempt to catch selected pokemon",
			callBack:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "inspect a given pokemon that you have captured",
			callBack:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show your pokedex of captured pokemon",
			callBack:    commandPokedex,
		},
	}
}

func commandHelp(cfg *config, _ string) error {
	commands := getCommands()

	fmt.Println()
	fmt.Printf("Welcome to %v! These are the available commands:\n", cliName)
	fmt.Println()

	// for future refernce: you can do for key, val := range getCommands() instead!
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandExit(cfg *config, _ string) error {
	os.Exit(0)
	return nil
}

func commandClear(cfg *config, _ string) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func commandExplore(cfg *config, areaName string) error {
	areaInfoResp, err := cfg.pokeapiClient.ListAreaInfo(areaName)
	if err != nil {
		return err
	}

	pokeEncounters := areaInfoResp.PokemonEncounters
	for _, encounter := range pokeEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, pokemonName string) error {
	pokemonInfoResp, err := cfg.pokeapiClient.PokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	// values for capture rate calc
	threshHold := 0.5
	C := 100.0
	baseExperience := float64(pokemonInfoResp.BaseExperience)
	randomNum := rand.Float64()

	captureRate := (randomNum / baseExperience) * C

	cfg.player.mu.Lock()
	defer cfg.player.mu.Unlock()

	_, ok := cfg.player.pokedex[pokemonName]
	if !ok {
		if captureRate >= threshHold {
			cfg.player.pokedex[pokemonName] = pokemonInfoResp
			fmt.Printf("Yay! %s was caught!\n", pokemonName)
		} else {
			fmt.Println("Darn! It got away!")
		}
	} else {
		fmt.Printf("You already posess a %s! Unable to capture another!\n", pokemonName)
	}

	return nil
}

func commandInspect(cfg *config, pokemonName string) error {
	pokemonInfo, ok := cfg.player.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("no pokemon in your pokedex that match the name %s", pokemonName)
	}

	name := pokemonInfo.Name
	height := pokemonInfo.Height
	weight := pokemonInfo.Weight
	stats := pokemonInfo.Stats
	types := pokemonInfo.Types

	fmt.Println()
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", name, height, weight)
	fmt.Printf("Stats:\n")
	for _, stat := range stats {
		fmt.Printf("    - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, ptype := range types {
		fmt.Printf("    - %s\n", ptype.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, _ string) error {
	if len(cfg.player.pokedex) < 1 {
		return fmt.Errorf("you have no pokemon in your pokedex. go capture some to add them")
	}

	fmt.Printf("Your Pokedex:\n")
	for key := range cfg.player.pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
