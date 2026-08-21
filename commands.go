package main

import (
	"fmt"
	"os"
	"github.com/Ceramik9/pokedex/internal/pokeapi"
)
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


func commandExplore(*config) error {

	return nil
}








