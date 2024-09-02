package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "Pokedex"

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

func formatInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startREPL(cfg *config) {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)

	printPrompt()
	for reader.Scan() {
		text := reader.Text()
		words := formatInput(text)

		if len(words) == 0 {
			continue
		}

		argument1 := ""

		if len(words) > 1 {
			argument1 = words[1]

		}

		commandName := words[0]

		if command, exists := commands[commandName]; exists {
			if err := command.callBack(cfg, argument1); err != nil {
				fmt.Println(err)
			}
		} else {
			printUnknown(text)
		}

		printPrompt()
	}
	fmt.Println()
}
