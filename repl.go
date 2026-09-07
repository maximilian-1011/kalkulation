package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Kalkulation >")
		reader.Scan()
		err := reader.Err()
		if err != nil {
			fmt.Println(err)
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unbekannter Befehl")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Split(output, " ")
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"bezugskalkulation": {
			name:        "bezugskalkulation",
			description: "Kombiniert Zieleinkaufspreis und Bareinkaufspreis mit bezugskosten",
			callback:    commandBezugskalkulation,
		},
		"zieleinkaufspreis": {
			name:        "zieleinkaufspreis",
			description: "Listenpreis nach abzug von Rabat",
			callback:    commandZieleinkaufspreis,
		},
		"bareinkaufspreis": {
			name:        "bareinkaufspreis",
			description: "Zieleinkaufspreis nach abzug von Skonto",
			callback:    commandBareinkaufspreis,
		},
		"hilfe": {
			name:        "hilfe",
			description: "Zeigt alle möglichen Befehle an",
			callback:    commandHelp,
		},
	}
}
