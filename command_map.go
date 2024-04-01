package main

import (
	"errors"
	"fmt"
)

func mapfCommand(cfg *config, args ...string) error {
		resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}

		cfg.nextLocationsURL = resp.Next
		cfg.prevLocationsURL = resp.Previous

		for _, location := range resp.Results {
			fmt.Println(location)
		}
		return nil

}

func mapbCommand(cfg *config, args ...string) error {
		if cfg.prevLocationsURL == nil {
			return errors.New("already on first page")
		}

		resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}
		cfg.nextLocationsURL = resp.Next
		cfg.prevLocationsURL = resp.Previous

		for _, location := range resp.Results {
			fmt.Println(location)
		}
		return nil
		
}