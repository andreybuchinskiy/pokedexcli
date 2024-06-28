package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	_, ok := cfg.caughtPokemon[pokemonName]
	if ok {
		return fmt.Errorf("pokemon already caught")
	}

	resp, err := cfg.client.GetPokemon(pokemonName)

	if err != nil {
		log.Fatal(err)
	}

	rand.New(rand.NewSource(99))
	maxBaseExperience := 300
	catchProbability := 1.0 - float64(resp.BaseExperience)/float64(maxBaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	randNumber := rand.Float64()
	if randNumber > catchProbability {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		cfg.caughtPokemon[pokemonName] = resp
		fmt.Printf("%s was caught!\n", pokemonName)
	}
	return nil
}
