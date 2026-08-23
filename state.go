package main

import (
	"github.com/Ceramik9/pokedex/internal/pokeapi"
)

type states struct {
	mapLocations  pokeapi.LocationArea
	pokemonList   pokeapi.LocationInfo
	pokemon       pokeapi.PokemonInfo
	caughtPokemon map[string]pokeapi.PokemonInfo
}

var state = states{
	caughtPokemon: make(map[string]pokeapi.PokemonInfo),
}
