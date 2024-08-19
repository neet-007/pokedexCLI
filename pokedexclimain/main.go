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
		Client: pokeapi.NewClient(time.Hour),
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

		commands := getCommands()

		command, ok := commands[text[0]]
		if !ok {
			continue
		}

		err := command.callback(&confing)
		if err != nil {
			fmt.Printf("error %s\n", err)
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
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
		"exploer": {
			name: "explore",
			description: "explore",
			callback: exploreCommand,
		}
	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}
