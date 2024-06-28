package main

import (
	"time"

	"git01.infra.andreybuchinskiy.com/andreybuchinskiy/pokedexcli/internal/pokeapi"
)

type config struct {
	client        pokeapi.Client
	nextURL       *string
	prevURL       *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		client:        pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
