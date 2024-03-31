package main

func getCommands() map[string]command {
	commandMap := map[string]command{
		"help": {
			name:        "help",
			description: "Displays helpful information",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    exitCommand,
		},
		"mapf": {
			name:        "map",
			description: "Displays next 20 locations from the Pokemon API",
			callback:    mapfCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations from the Pokemon API",
			callback:    mapbCommand,
		},
		"explore": {
			name: "explore",
			description: "Explore an area for Pokemopn and items",
			callback: exploreCommand,
		},
	}

	return commandMap
}