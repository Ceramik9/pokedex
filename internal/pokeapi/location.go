package pokeapi

import (
	"io"
	"fmt"
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []results `json:"results"`
}


func (la *LocationArea) Update(url string) error {
	
	//check for cached data
	data, ok := pokeCache.Get(url)
	if ok {
		err := json.Unmarshal(data, la)
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

	// check response status code
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK HTTP status: %s\n", res.Status)
	}
	
	// save response as []byte1
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %w", err)
	}
	// cache data
	pokeCache.Add(url, data)

	// update data from response
	err = json.Unmarshal(data, la)
	if err != nil {
		return fmt.Errorf("Decoding failed: %w", err)
	}
	return nil
}

func (la LocationArea) PrintData() error {
	if la.Results == nil {
		return fmt.Errorf("locations are empty")
	}
	for _, location := range la.Results {
		fmt.Println(location.Name)
	}
	return nil
}


type results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


