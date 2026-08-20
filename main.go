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

		// user commnd input
		scanner.Scan()
		inputCommand := scanner.Text()

		// check if the command is a valid command
		command, ok := commandList.registry[inputCommand]
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

