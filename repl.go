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
		fmt.Print("Kalkulation > ")
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
			fmt.Println()
			fmt.Println("Unbekannter Befehl")
			fmt.Println("Gib 'help' ein für wietere Informationen")
			fmt.Println()
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
		"help": {
			name:        "help",
			description: "Zeigt alle möglichen Befehle an",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Schließt das Programm",
			callback:    commandExit,
		},
		"bk": {
			name:        "bk",
			description: "Vollständige Bezugskalkulation",
			callback:    commandBezugskalkulation,
		},
		"zek": {
			name:        "zek",
			description: "Zieleinkaufspreis. Listenpreis nach abzug von Rabatt",
			callback:    commandZieleinkaufspreis,
		},
		"bek": {
			name:        "bek",
			description: "Bareinkaufspreis. Zieleinkaufspreis nach abzug von Skonto",
			callback:    commandBareinkaufspreis,
		},
		"us": {
			name:        "us",
			description: "Umschlagshäufigkeit",
			callback:    commandUmschlagshäufigkeit,
		},
		"ld": {
			name:        "ld",
			description: "Lagerdauer",
			callback:    commandLagerdauer,
		},
		"dsb": {
			name:        "dsb",
			description: "Durchschnittlicher Lagerbestand",
			callback:    commandDurchschnittlicherLagerbestand,
		},
		"lzs": {
			name:        "lzs",
			description: "Lagerzinssatz",
			callback:    commandLagerzinssatz,
		},
		"lz": {
			name:        "lz",
			description: "Lagerzins",
			callback:    commandLagerzins,
		},
		"skp": {
			name:        "skp",
			description: "Selbstkostenpreis",
			callback:    commandSelbstkostenpreis,
		},
		"nvp": {
			name:        "nvp",
			description: "Nettoverkaufspreis",
			callback:    commandNettoVerkaufsPreis,
		},
		"bvp": {
			name:        "bvp",
			description: "Bruttoverkaufspreis",
			callback:    commandBruttoVerkaufsPreis,
		},
	}
}
