package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	locationsConfig := Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := scanner.Text()
		inputWords := cleanInput(words)

		if len(inputWords) == 0 {
			continue
		}

		commandName := inputWords[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(&locationsConfig)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(strings.ToLower(text))
	return cleanedInput
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display first 20 location areas.\nSubsequent calls display next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 location areas.",
			callback:    commandMapB,
		},
	}
}

type Config struct {
	Next     string
	Previous string
}
