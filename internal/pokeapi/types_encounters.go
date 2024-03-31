package pokeapi

type PokemonInfo struct {
	Name string `json:"name"`
}

type PokemonEncounter struct {
	Pokemon PokemonInfo `json:"pokemon"`
}

type AreaAPIResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}