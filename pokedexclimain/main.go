package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/neet-007/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	confing := pokeapi.Config{
		Client:     pokeapi.NewClient(time.Hour),
		PokemonMap: map[string]pokeapi.PokemonResponse{},
	}
	for {
		fmt.Print("pokedex >")

		status := scanner.Scan()
		if !status {
			fmt.Println("an error happend while reading input please try again")
			continue
		}
		text := cleanInput(scanner.Text())

		if len(text) == 0 {
			continue
		}

		args := []string{}

		if len(text) > 1 {
			args = text[1:]
		}
		commands := getCommands()

		command, ok := commands[text[0]]
		if !ok {
			continue
		}

		err := command.callback(&confing, args...)
		if err != nil {
			fmt.Printf("error %s\n", err)
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: " Exit the Pokedex",
			callback:    exitCommand,
		},
		"map": {

			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world",
			callback:    mapCommand,
		},
		"mapb": {

			name:        "map",
			description: "Similar to the map command, however, instead of displaying the next 20 locations,",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore",
			description: "explore {loaction name}",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "catch {pokemon name}",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "pokedex",
			callback:    pokedexCommand,
		},
	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}
