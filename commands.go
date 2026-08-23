package main

import (
	"fmt"
	"errors"
	"strings"
	"os"
	"github.com/Ceramik9/pokedex/internal/pokeapi"
)


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
	fmt.Println("explore: Displays a list of Pokempns in area")
	fmt.Println("catch: Attempt to catch the Pokemon")
	return nil
}


func commandMap(*config) error {
	
	//set url
	var locationAreaURL string
	if state.mapLocations.Next != "" {
		locationAreaURL = state.mapLocations.Next
	} else {
		locationAreaURL = pokeapi.DefaultMapURL
	}

	//update locations
	state.mapLocations.Update(locationAreaURL)

	//print list of locations
	state.mapLocations.PrintData()

	return nil
}


func commandMapb(*config) error {
	
	// checks for first and last page numbers
	if state.mapLocations.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	if pokeapi.LastPrevious == state.mapLocations.Previous && pokeapi.FirstNext == state.mapLocations.Next {
		fmt.Println("you're on the first page")
		return nil
	}
	
	//update locations
	state.mapLocations.Update(state.mapLocations.Previous)

	// print list of mapLocations
	state.mapLocations.PrintData()
	
	return nil
}


func commandExplore(*config) error {
	fmt.Printf("Exploring %s...\n", userCommand.arg)

	// check for errors
	err := state.pokemonList.Update(pokeapi.DefaultMapURL, userCommand.arg)
	if err != nil {
		return errors.New("invalid location")
	}
	// print pokemon list
	state.pokemonList.PrintData()

	return nil
}


func commandCatch(*config) error {
	
	// return error if pokemon name is missing
	if userCommand.arg == "" {
		return errors.New("you must provide Pokemon name")
	}
	
	// return error if pokemon name does not exist
	err := state.pokemon.Update(pokeapi.DefaultPokemonURL, userCommand.arg)
	if err != nil {
		return fmt.Errorf("Pokemon %s does not exist", userCommand.arg)
	}
	
	// attemp to catch pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", state.pokemon.Name)
	result := state.pokemon.Catch(state.caughtPokemon)
	if !result {
		fmt.Printf("%s escaped!\n", state.pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", state.pokemon.Name)
	return nil
}


