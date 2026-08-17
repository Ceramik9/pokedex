package main

import (
	"strings"
	"fmt"
	"os"
)


func cleanInput(text string) []string {
	if len(text) == 0 || text == " " {
		return []string{}
	}
	return strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}


