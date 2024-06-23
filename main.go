package main

import (
	"time"

	"git01.infra.andreybuchinskiy.com/andreybuchinskiy/pokedexcli/internal/pokeapi"
)

type config struct {
	client  pokeapi.Client
	nextURL *string
	prevURL *string
}

func main() {
	cfg := config{
		client: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
