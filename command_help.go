package main

import "fmt"

func helpCommand(cfg *config, args ...string) error {
		fmt.Println()
		fmt.Println("Welcome to your Pokedex!")
		fmt.Println("Commands:")
		fmt.Println()
		for _, command := range getCommands() {
			fmt.Printf("%s: %s\n", command.name, command.description)
		}
		fmt.Println()
		return nil
	}