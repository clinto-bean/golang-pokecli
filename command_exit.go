package main

import (
	"fmt"
	"os"
	"strings"
)

func exitCommand(cfg *config, param string) error {
		fmt.Println("Are you sure you want to exit? (yes/y or no/n)")
		scanner.Scan()
		text := strings.ToLower(scanner.Text())
		if text == "yes" || text == "y"{
			fmt.Println("Exiting, thank you for using PokeCLI!")
			os.Exit(0)
		} else {
			fmt.Println("Returning to the main menu.")
		}
	return nil
	}