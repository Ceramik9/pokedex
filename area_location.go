package main

type locationArea struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []results `json:"results"`
}

func (l locationArea) next() {
	
}

type results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// allocates locationArea in memory
var mapLocations locationArea

// default map url
const defaultMapURL = "https://pokeapi.co/api/v2/location-area"

//used to check if the user is already on the first page
const lastPrevious = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
const firstNext = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

//used to check if the user is already on the last page
const lastNext = "https://pokeapi.co/api/v2/location-area?offset=1520&limit=20"
