package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/neet-007/pokeapi"
)

func catchCommand(config *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon name")
	}
	pokemon, err := config.Client.CatchPokemon(args[0])

	if err != nil {
		return err
	}

	threshold := 50
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("was not couch %v\n", args[0])
	}

	config.PokemonMap[args[0]] = pokemon
	fmt.Printf("was cought %v\n", args[0])

	return nil
}
