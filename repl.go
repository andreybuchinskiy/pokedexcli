package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "Pokedex > "

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print(prompt)
		scanner.Scan()
		cmd := scanner.Text()
		commandName, ok := commands[cmd]
		if !ok {
			fmt.Println("Invalid Command!")
			continue
		}
		err := commandName.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
