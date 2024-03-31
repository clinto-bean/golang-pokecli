package main

import "fmt"

func helpCommand(cfg *config, param string) error {
		fmt.Println("Available commands:")
		for _, command := range getCommands() {
			fmt.Printf("%s: %s\n", command.name, command.description)
		}
		return nil
	}