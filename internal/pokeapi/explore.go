package pokeapi

import (
	"fmt"
	"io"
	"time"
	"net/http"
	"encoding/json"
	"github.com/Ceramik9/pokedex/internal/pokecache"
)

// create a new struct for cashes
// initialise pokeCashe
var pokeCache = pokecache.NewCache(5 * time.Second)

func (li *LocationInfo) Update(url, location string) error {
	
	fullURL := url + "/" + location

	//check for cached data
	data, ok := pokeCache.Get(fullURL)
	if ok {
		err := json.Unmarshal(data, li)
		if err != nil {
			return fmt.Errorf("unmarshal cache data failed: %w", err)
		}
		return nil
	}

	//request data
	res, err := http.Get(fullURL)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()
	
	// check response status code
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK HTTP status: %s\n", res.Status)
	}

	// save response as []byte1
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	// cache data
	pokeCache.Add(fullURL, data)

	// update data from response
	err = json.Unmarshal(data, li)
	if err != nil {
		return fmt.Errorf("decoding failed: %w", err)
	}
	return nil
}


func (li LocationInfo) PrintData() error {

	// check if any pokemon have been found in the area
	if len(li.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found")
		return nil
	}
	// Display list of pokemon in the area
	fmt.Println("Found Pokemon:")
	for _, encounter := range li.PokemonEncounters {
		pokemon := encounter.Pokemon.Name
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}


