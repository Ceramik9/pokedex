package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputCommand := scanner.Text()
		command, ok := supportedCommands[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback()
		}
	}
}

