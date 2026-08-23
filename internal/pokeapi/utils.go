package pokeapi

import (
	"time"
	"github.com/Ceramik9/pokedex/internal/pokecache"
)

// cache for storing api responses
var pokeCache = pokecache.NewCache(5 * time.Second)

// default pokemon url
const DefaultPokemonURL = "https://pokeapi.co/api/v2/pokemon"

// default map url
const DefaultMapURL = "https://pokeapi.co/api/v2/location-area"

//used to check if the user is already on the first page
const LastPrevious = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
const FirstNext = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

//used to check if the user is already on the last page
const LastNext = "https://pokeapi.co/api/v2/location-area?offset=1520&limit=20"
