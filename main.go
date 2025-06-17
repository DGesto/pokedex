package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(strings.ToLower(text))
	return cleanedInput
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		s := scanner.Text()
		inputWords := cleanInput(s)
		if val, ok := myCommands[inputWords[0]]; ok {
			err := val.callback()
			if err != nil {
				fmt.Println("An error has occurred: %w", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
