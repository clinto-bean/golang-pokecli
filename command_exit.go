package main

import (
	"fmt"
	"os"
)

func exitCommand(cfg *config, args ...string) error {
		fmt.Println("Exiting, thank you for using PokeCLI!")
		os.Exit(0)
		return nil
	}