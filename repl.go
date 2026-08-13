package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func startRepl() {
	scanner := bufio.NewReader(os.Stdin)
	state := &config{commands: getCommand()}
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
		result, ok := getCommand()[cmd]

		if !ok {
			fmt.Printf("Unknown command: %s\n", cmd)
			result = getCommand()["help"]
		}

		err = result.callback(state)

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
}

type cliCommand struct {
	name        string
	description string
	callback    func(state *config) error
}

func getCommand() map[string] cliCommand{
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
	}		
}
