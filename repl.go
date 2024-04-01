package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/clinto-bean/golang-pokecli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
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
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type command struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    catchCommand,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    exploreCommand,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    mapfCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    mapbCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
		"inspect": {
			name: "inspect",
			description: "Display stats of caught pokemon by name",
			callback: inspectCommand,
		},
	}
}