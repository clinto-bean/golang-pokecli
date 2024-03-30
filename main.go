package main

import (
	"time"

	"github.com/clinto-bean/golang-pokecli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.CreateClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)	
}