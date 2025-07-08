package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
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
		"map": {
			Name:        "map",
			Description: "Display the names of 20 location areas",
			Callback:    CommandMap,
		},
	}
}
