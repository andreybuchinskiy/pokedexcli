package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.client.ListLocationAreas(cfg.nextURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas")
	for _, area := range resp.Results {
		fmt.Printf(" -- %s\n", area.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.nextURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.client.ListLocationAreas(cfg.prevURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas")
	for _, area := range resp.Results {
		fmt.Printf(" -- %s\n", area.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous
	return nil
}
