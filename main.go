package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		result, ok := getCommand()[cmd]

		if !ok {
			fmt.Printf("Unknown command: %s\n", cmd)
			result = getCommand()["help"]
		}

		err = result.callback()

		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
	}
}
