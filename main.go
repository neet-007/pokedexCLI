package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

		command.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}
