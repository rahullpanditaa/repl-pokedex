package commands

import (
	"fmt"

	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

func CommandHelp(*utils.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, cmd := range CommandsRegistry {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}
	return nil
}
