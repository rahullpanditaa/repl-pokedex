package commands

import (
	"fmt"

	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

func CommandHelp(config *utils.Config, areaName string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, cmd := range CommandsRegistry {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}
	return nil
}
