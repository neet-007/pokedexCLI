package main

import (
	"errors"
	"fmt"

	"github.com/neet-007/pokeapi"
)

func exploreCommand(config *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a location")
	}
	location, err := config.Client.ListLocation(args[0])

	if err != nil {
		return err
	}

	fmt.Println("the pokemons in location %s", location.Name)
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("-%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
