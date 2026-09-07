package main

import "fmt"

func commandHelp() error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("Willkommen im Kalkulationsprogramm!")
	fmt.Println("Nutzung:")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
