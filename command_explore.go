package main

import (
	"fmt"
)

func exploreCommand(cfg *config, location string) error {
	resp, err := cfg.pokeapiClient.ExploreLocation(location)

	if err != nil {
		return err
	}
	fmt.Println("Found pokemon:")
	var names []string
	for _, encounter := range resp.PokemonEncounters {
		names = append(names, encounter.Pokemon.Name)
	}

	for _, name := range names {
		fmt.Printf("- %v\n", name)
	}

	return nil
}