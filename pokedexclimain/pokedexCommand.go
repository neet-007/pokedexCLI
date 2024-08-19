package main

import (
	"fmt"

	"github.com/neet-007/pokeapi"
)

func pokedexCommand(config *pokeapi.Config, args ...string) error {
	for _, pokemon := range config.PokemonMap {
		fmt.Printf("%s\n", pokemon.Name)
	}

	return nil
}
