package main

import (
	"errors"
	"fmt"

	"github.com/neet-007/pokeapi"
)

func inspectCommand(config *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon name")
	}

	pokemon, ok := config.PokemonMap[args[0]]
	if !ok {
		return errors.New("you havent cought this pokemon")
	}

	fmt.Printf("Name:%s\n", pokemon.Name)
	fmt.Printf("Forms:%v\n", pokemon.Forms)
	fmt.Printf("Hegiht:%v\n", pokemon.Height)
	fmt.Printf("Wegiht:%v\n", pokemon.Weight)

	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s:%v\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, type_ := range pokemon.Types {
		fmt.Printf("-%s:%v\n", type_.Type.Name, type_.Slot)
	}
	return nil
}
