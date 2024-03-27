package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type command struct {
	name string
	description string
	callback func() error
}

func helpCommand() error {
		fmt.Println("Available commands:")
		for _, command := range initCommands() {
			fmt.Printf("%s: %s\n", command.name, command.description)
		}
		return nil
	}

func exitCommand() error {
		fmt.Println("Are you sure you want to exit? (yes/y or no/n)")
		scanner.Scan()
		text := strings.ToLower(scanner.Text())
		if text == "yes" || text == "y"{
			fmt.Println("Exiting, thank you for using PokeCLI!")
			os.Exit(0)
			return nil
		} else {
			fmt.Println("Returning to the main menu.")
		}
		
		
		
		return nil
	}

func mapCommand() error {
		fmt.Println("Map command")
		return nil
	}

func mapbCommand() error {
		fmt.Println("Mapb command")
		return nil
	}

func initCommands() map[string]command {
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

func startRepl() map[string]command {

	commands := initCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan(){
			text := scanner.Text()
			if command, found := commands[text]; found {
				command.callback()
			} else {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
			}
		}
	}
}