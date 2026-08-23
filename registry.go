package main


type userInput struct {
	command string
	arg     string
}

var userCommand userInput

type config struct {
	registry map[string]cliCommand
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

// list of valid commands
var commandList = config{
	registry: map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays next 20 locations in Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous 20 locations in Pokemon world",
			callback: commandMapb,
		},
		"explore": {
		name: "explore",
		description: "Displays a list of Pokempns in area",
		callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch the Pokemon",
			callback: commandCatch,
		},
	},
}


