package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rahullpanditaa/repl-pokedex/repl"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := repl.CleanInput(input)
		firstWord := cleanedInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}

}
