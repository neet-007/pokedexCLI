package main

import (
	"fmt"
	"github.com/neet-007/pokeapi"
	"log"
)

func mapCommand(config *pokeapi.Config, args ...string) error {
	locations, err := config.Client.ListLocations(config.Next)

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
