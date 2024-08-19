package main

import (
	"errors"
	"fmt"
	"github.com/neet-007/pokeapi"
	"log"
)

func mapbCommand(config *pokeapi.Config) error {
	if config.Previous == nil {
		return errors.New("you are in the first page")
	}
	locations, err := config.Client.ListLocations(config.Previous)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("the list of locations")
	for _, location := range locations.Results {
		fmt.Printf("-%s\n", location.Name)
	}

	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}
