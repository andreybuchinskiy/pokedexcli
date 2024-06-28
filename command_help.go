package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	commands := getCommands()
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Printf("\n")
	return nil
}
