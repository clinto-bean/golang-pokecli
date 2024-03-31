package main

import (
	"time"

	"github.com/clinto-bean/golang-pokecli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	
}

type command struct {
	name        string
	description string
	callback    func(cfg *config, arg string) error
	
}

func main() {
	pokeClient := pokeapi.CreateClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)	
}