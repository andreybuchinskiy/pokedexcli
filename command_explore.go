package main

import (
	"errors"
	"fmt"
	"log"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location provided")
	}
	location := args[0]
	resp, err := cfg.client.GetLocationAreas(location)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pokemon in area: ", resp.Name)
	for _, area := range resp.PokemonEncounters {
		fmt.Printf(" -- %s\n", area.Pokemon.Name)
	}
	return nil
}
