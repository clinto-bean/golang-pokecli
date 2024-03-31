package main

import (
	"errors"
	"fmt"
)

func mapfCommand(cfg *config, param string) error {
		resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}

		cfg.nextLocationsURL = &resp.Next
		cfg.prevLocationsURL = &resp.Previous

		for _, location := range resp.Results {
			fmt.Println(location)
		}
		return nil

}

func mapbCommand(cfg *config, param string) error {
		if cfg.prevLocationsURL == nil {
			return errors.New("you're on the first page")
		}

		resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}
		cfg.nextLocationsURL = &resp.Next
		cfg.prevLocationsURL = &resp.Previous

		for _, location := range resp.Results {
			fmt.Println(location)
		}
		return nil
		
}