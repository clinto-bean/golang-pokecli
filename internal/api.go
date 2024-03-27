package main

import (
	"io"
	"log"
	"net/http"
)

type config struct {
	Next string
	Previous string
}

func getAPIConfig() config {
	return config{
		Next: "https://pokeapi.co/api/v2/pokemon?offset=20&limit=20",
		Previous: "",
	}
}

func getPokemon() {
	apiConfig := getAPIConfig()
	
	res, err := http.Get(apiConfig.Next)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}


}