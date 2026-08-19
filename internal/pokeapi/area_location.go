package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type locationArea struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []results `json:"results"`
}

func (l locationArea) Update(url string) error {
	
	//request data
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Request failed: %w", err)
	}
	defer res.Body.Close()
	
	//decode data and update mapLocations
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&MapLocations)
	if err != nil {
		return fmt.Errorf("Decoding failed: %w", err)
	}
	return nil
}

func (l locationArea) PrintLocation() error {
	if MapLocations.Results == nil {
		return fmt.Errorf("locations are empty")
	}
	for _, location := range MapLocations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

type results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// allocates locationArea in memory
var MapLocations locationArea

// default map url
const DefaultMapURL = "https://pokeapi.co/api/v2/location-area"

//used to check if the user is already on the first page
const LastPrevious = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
const FirstNext = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

//used to check if the user is already on the last page
const LastNext = "https://pokeapi.co/api/v2/location-area?offset=1520&limit=20"
