package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	// initialise io buffer
	scanner := bufio.NewScanner(os.Stdin)

	// program loop
	for {
		fmt.Print("Pokedex > ")

		// user command input
		scanner.Scan()
		inputCommand := scanner.Text()

		// parse user command input
		userCommand.command, userCommand.arg = commandParser(inputCommand)

		// check if the command is a valid command
		command, ok := commandList.registry[userCommand.command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&commandList)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

