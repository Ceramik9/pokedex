package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"math/rand"
	"encoding/json"
)

func (pi *PokemonInfo) Update(url, location string) error {
	
	fullURL := url + "/" + location

	//check for cached data
	data, ok := pokeCache.Get(fullURL)
	if ok {
		err := json.Unmarshal(data, pi)
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
	err = json.Unmarshal(data, pi)
	if err != nil {
		return fmt.Errorf("decoding failed: %w", err)
	}
	return nil
}


func (pi PokemonInfo) Catch(pokemonList map[string]PokemonInfo) bool {
	
	//set chance
	var chance int
	if pi.BaseExperience < 100 {
		chance = 80
	} else if pi.BaseExperience < 200 {
		chance = 60
	} else {
		chance = 40
	}

	// catch attempt
	attempt := (rand.Intn(pi.BaseExperience) * 100) / pi.BaseExperience
	
	//result
	if attempt >= chance {
		pokemonList[pi.Name] = pi
		return true
	}
	return false
}







