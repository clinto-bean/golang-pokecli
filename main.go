package main

import (
	"time"

	"github.com/clinto-bean/golang-pokecli/internal/pokeapi"
)


func main() {
	serverTimeout := time.Second * 5
	cacheInterval := time.Minute * 5
	pokeClient := pokeapi.CreateClient(serverTimeout, cacheInterval)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)	
}