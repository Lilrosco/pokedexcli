package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()

			if len(input) == 0 {
				continue
			}

			cleanInput := cleanInput(input)
			fmt.Printf("Your command was: %s\n", cleanInput[0])
		}
	}
}
