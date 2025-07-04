package repl

import "fmt"

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, cmd := range CommandsRegistry {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}
	return nil
}
