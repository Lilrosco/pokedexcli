package main

import (
	"fmt"
)

func commandHelp(state *config) error {
	fmt.Print("Usage: \n\n")

	for _, cmd := range getCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
