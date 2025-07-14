package commands

import "github.com/rahullpanditaa/repl-pokedex/internal/utils"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*utils.Config) error
}

var CommandsRegistry map[string]CliCommand

func init() {
	CommandsRegistry = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"mapf": {
			Name:        "mapf",
			Description: "Display the names of 20 location areas",
			Callback:    CommandMapForward,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Go back a page",
			Callback:    CommandMapBack,
		},
	}
}
