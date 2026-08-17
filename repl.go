package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func startRepl(cfg *config) {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	for {
		fmt.Print("Pokedex > ")
		input, err := scanner.ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured: %v", err)
		}

		if len(input) == 0 {
			continue
		}

		cmd := cleanInput(input)[0]
		result, ok := cfg.commands[cmd]

		if !ok {
			fmt.Printf("Unknown command: %s\n", cmd)
			result = cfg.commands["help"]
		}

		err = result.callback(cfg)

		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type config struct {
	commands map[string] cliCommand
	baseLocationAreasURL string
	nextLocationAreasURL string
	prevLocationAreasURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string] cliCommand{
	return map[string] cliCommand {
		"exit": {
        	name:        "exit",
        	description: "Exit the Pokedex",
        	callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 location areas",
			callback:    commandMapb,
		},
	}		
}
