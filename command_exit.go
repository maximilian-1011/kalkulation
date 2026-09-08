package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println()
	fmt.Println("Schließe Kalkulationsprogramm...")
	fmt.Println("Auf Wiedersehen!")
	fmt.Println()
	os.Exit(0)
	return nil
}
