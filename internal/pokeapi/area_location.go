package pokeapi

import (
	"io"
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"github.com/Ceramik9/pokedex/internal/pokecache"
)

type locationArea struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []results `json:"results"`
}

var pokeCache = pokecache.NewCache(5 * time.Second)

func (l *locationArea) Update(url string) error {
	
	//check for cached data
	data, ok := pokeCache.Get(url)
	if ok {
		err := json.Unmarshal(data, &l)
		if err != nil {
			return fmt.Errorf("Unmarshal cache data failer: %w", err)
		}
		return nil
	}

	//request data
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Request failed: %w", err)
	}
	defer res.Body.Close()
	
	// save response as []byte
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %w", err)
	}
	// cache data
	pokeCache.Add(url, data)

	// update data from response
	err = json.Unmarshal(data, &l)
	if err != nil {
		return fmt.Errorf("Decoding failed: %w", err)
	}
	return nil
}

func (l locationArea) PrintLocation() error {
	if MapLocations.Results == nil {
		return fmt.Errorf("locations are empty")
	}
	for _, location := range l.Results {
		fmt.Println(location.Name)
	}
	return nil
}

type results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// allocate locationArea in memory
var MapLocations locationArea

// default map url
const DefaultMapURL = "https://pokeapi.co/api/v2/location-area"

//used to check if the user is already on the first page
const LastPrevious = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
const FirstNext = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

//used to check if the user is already on the last page
const LastNext = "https://pokeapi.co/api/v2/location-area?offset=1520&limit=20"
