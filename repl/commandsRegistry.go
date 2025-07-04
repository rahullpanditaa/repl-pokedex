package repl

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var CommandsRegistry = map[string]CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
}
