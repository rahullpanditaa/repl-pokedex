package commands

import (
	"fmt"
	"os"

	"github.com/rahullpanditaa/repl-pokedex/internal/utils"
)

func CommandExit(config *utils.Config, areaName string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
