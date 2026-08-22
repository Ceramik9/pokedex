package main

import (
	"fmt"
	"errors"
	"strings"
	"os"
	"github.com/Ceramik9/pokedex/internal/pokeapi"
)

var command string
var argument string


func cleanInput(text string) []string {
	if len(text) == 0 || text == " " {
		return []string{}
	}
	return strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
}


func commandParser(input string) (string, string) {
	command :=  cleanInput(input)
	if len(command) > 2 {
		fmt.Println("Too many argumants")
	}
	if len(command) == 2 {
		return command[0], command[1]
	}
	if len(command) == 1 {
		return command[0], ""
	}
	return "", ""
}



// initialise LocationArea struct used in map and mapb
var MapLocations pokeapi.LocationArea


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
	if MapLocations.Next != "" {
		locationAreaURL = MapLocations.Next
	} else {
		locationAreaURL = pokeapi.DefaultMapURL
	}

	//update locations
	MapLocations.Update(locationAreaURL)

	//print list of locations
	MapLocations.PrintLocation()

	return nil
}


func commandMapb(*config) error {
	
	// checks for first and last page numbers
	if MapLocations.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	if pokeapi.LastPrevious == MapLocations.Previous && pokeapi.FirstNext == MapLocations.Next {
		fmt.Println("you're on the first page")
		return nil
	}
	
	//update locations
	MapLocations.Update(MapLocations.Previous)

	// print list of MapLocations
	MapLocations.PrintLocation()
	
	return nil
}

// move this line under a struct with MapLocations
// initialise LocationInfo
var pokemonList pokeapi.LocationInfo

func commandExplore(*config) error {
	fmt.Printf("Exploring %s...\n", userCommand.arg)

	// check for errors
	err := pokemonList.Update(pokeapi.DefaultMapURL, userCommand.arg)
	if err != nil {
		return errors.New("Invalid location")
	}
	// print pokemon list
	pokemonList.PrintData()

	return nil
}


