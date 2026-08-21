package pokeapi

import (
	"fmt"
	"io"
	"time"
	"net/http"
	"encoding/json"
	"github.com/Ceramik9/pokedex/internal/pokecache"
)

// initialise pokeCashe
var pokeCache = pokecache.NewCache(5 * time.Second)

func (li LocationInfo) Update(url, location string) error {
	
	fullURL := url + "/" + location

	//check for cached data
	data, ok := pokeCache.Get(fullURL)
	if ok {
		err := json.Unmarshal(data, &li)
		if err != nil {
			return fmt.Errorf("Unmarshal cache data failer: %w", err)
		}
		return nil
	}

	//request data
	res, err := http.Get(fullURL)
	if err != nil {
		return fmt.Errorf("Request failed: %w", err)
	}
	defer res.Body.Close()

	// save response as []byte1
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %w", err)
	}
	// cache data
	pokeCache.Add(fullURL, data)

	// update data from response
	err = json.Unmarshal(data, &li)
	if err != nil {
		return fmt.Errorf("Decoding failed: %w", err)
	}
	return nil
}


func (li LocationInfo) PrintData() error {
	// check if there are any pokemons

	for pokemon, _ := range li.PokemonEncounters {
	fmt.Println(pokemon)
	}
	



	//if l.Results == nil {
	//	return fmt.Errorf("locations are empty")
	//}
	return nil
}





