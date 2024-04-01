package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchCommand(cfg *config, args ...string) error {
	if args[0] == "" {
		return errors.New("you must provide a pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught. It has been added to your Pokedex as entry %v.\n", pokemon.Name, pokemon.ID)

	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}