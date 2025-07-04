package main

import (
	// "github.com/rahullpanditaa/repl-pokedex/repl"
	"github.com/rahullpanditaa/repl-pokedex/cache"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	// 	fmt.Print("Pokedex > ")
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	input := scanner.Text()
	// 	cleanedInput := repl.CleanInput(input)
	// 	firstWord := cleanedInput[0]

	// 	if command, exists := repl.CommandsRegistry[firstWord]; exists {
	// 		err := command.Callback()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	} else {
	// 		fmt.Println("Unknown command")
	// 	}
	// }
	cache.MapCommand()

}
