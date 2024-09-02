package main

import (
	"fmt"
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
