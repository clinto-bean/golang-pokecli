package main

import (
	"errors"
	"fmt"
)

func pokedexCommand(cfg *config, args ...string) error {
	dex := cfg.pokedex
	if len(dex) == 0 {
		return errors.New("you haven't caught anything yet")
	}
	fmt.Println("Your pokedex includes:")
	for _, entry := range dex {
		fmt.Printf("%v - %v\n", entry.ID, entry.Name )
	}
	return nil
}