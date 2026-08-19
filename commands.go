package main

import (
	"strings"
	"fmt"
	"os"
	"github.com/Ceramik9/pokedex/internal/pokeapi"
)

// possibly not needed
// delete in the future if not used
func cleanInput(text string) []string {
	if len(text) == 0 || text == " " {
		return []string{}
	}
	return strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays next 20 locations in Pokemon world")
	fmt.Println("mapb: Displays previous 20 locations in Pokemon world")
	return nil
}

func commandMap(*config) error {
	
	//set url
	var locationAreaURL string
	if pokeapi.MapLocations.Next != "" {
		locationAreaURL = pokeapi.MapLocations.Next
	} else {
		locationAreaURL = pokeapi.DefaultMapURL
	}

	//update locations
	pokeapi.MapLocations.Update(locationAreaURL)

	//print list of locations
	pokeapi.MapLocations.PrintLocation()

	return nil
}

func commandMapb(*config) error {
	
	// checks for first and last page numbers
	if pokeapi.MapLocations.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	if pokeapi.LastPrevious == pokeapi.MapLocations.Previous && pokeapi.FirstNext == pokeapi.MapLocations.Next {
		fmt.Println("you're on the first page")
		return nil
	}
	
	//update locations
	pokeapi.MapLocations.Update(pokeapi.MapLocations.Previous)

	// print list of MapLocations
	pokeapi.MapLocations.PrintLocation()
	
	return nil
}


