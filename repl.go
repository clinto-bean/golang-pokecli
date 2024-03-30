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
}

var scanner = bufio.NewScanner(os.Stdin)

func startRepl(cfg *config) {
	

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := parseInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		
		command, exists := getCommands()[words[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

func parseInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type command struct {
	name string
	description string
	callback func(*config) error
}

func getCommands() map[string]command {
	commandMap := map[string]command{
			"help": {
				name: "help",
				description: "Displays helpful information",
				callback: helpCommand,
			},
			"exit": {
				name: "exit",
				description: "Exit the program",
				callback: exitCommand,
			},
			"map": {
				name: "map",
				description: "Displays next 20 locations from the Pokemon API",
				callback: mapCommand,
			},
			"mapb": {
				name: "mapb",
				description: "Displays previous 20 locations from the Pokemon API",
				callback: mapbCommand,
			},
		
	}
	
	return commandMap
}