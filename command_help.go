package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	// fmt.Println("Here are your available commands: ")

	availableCommands := getCommands()
	for _, cmd := range availableCommands{
		fmt.Printf(" - %s: %s", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}

