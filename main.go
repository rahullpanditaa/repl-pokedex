package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rahullpanditaa/repl-pokedex/internal/commands"
	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := utils.CleanInput(input)
		firstWord := cleanedInput[0]

		if command, exists := commands.CommandsRegistry[firstWord]; exists {
			err := command.Callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
