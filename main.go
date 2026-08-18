package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	
	con := config{
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
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputCommand := scanner.Text()
		command, ok := con.registry[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&con)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

