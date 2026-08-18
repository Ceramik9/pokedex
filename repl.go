package main

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
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
	if mapLocations.Next != "" {
		locationAreaURL = mapLocations.Next
	} else {
		locationAreaURL = defaultMapURL
	}

	//update locations
	updateLocations(locationAreaURL)

	//print list of locations
	printLocations()

	return nil
}

func commandMapb(*config) error {
	
	// check page number
	if mapLocations.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	
	if lastPrevious == mapLocations.Previous && firstNext == mapLocations.Next {
		fmt.Println("you're on the first page")
		return nil
	}
	
	//update locations
	updateLocations(mapLocations.Previous)

	// print list of locations
	printLocations()
	
	return nil
}

func updateLocations(url string) error {
	
	//request data
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Request failed: %w", err)
	}
	defer res.Body.Close()
	
	//decode data and update mapLocations
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&mapLocations)
	if err != nil {
		return fmt.Errorf("Decoding failed: %w", err)
	}
	return nil
}

func printLocations() error {
	if mapLocations.Results == nil {
		return fmt.Errorf("locations are empty")
	}
	for _, location := range mapLocations.Results {
		fmt.Println(location.Name)
	}
	//debug to delete
	//fmt.Printf("Next: %s\n", mapLocations.Next)
	//fmt.Printf("Previous: %s\n", mapLocations.Previous)
	return nil
}













