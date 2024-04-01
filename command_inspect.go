package main

import (
	"errors"
	"fmt"
)

func inspectCommand(cfg *config, args ...string) error {
	dex := cfg.pokedex
	name := args[0]
	pokemon := dex[name]
	if pokemon.Name == "" {
		return errors.New("you have not caught this pokemon")
	}
	fmt.Println()
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Stats:")
		fmt.Println(" -Height:", pokemon.Height)
		fmt.Println(" -Weight:", pokemon.Weight)
		for _, s := range pokemon.Stats {
			fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Println(" -", t.Type.Name)
		}
		fmt.Println()

	
	return nil
}