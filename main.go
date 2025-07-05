package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rahullpanditaa/repl-pokedex/repl"
)

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := repl.CleanInput(input)
		firstWord := cleanedInput[0]

		if command, exists := repl.CommandsRegistry[firstWord]; exists {
			err := command.Callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
