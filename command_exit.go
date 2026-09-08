package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Schließe Kalkulationsprogramm...")
	fmt.Println("Auf Wiedersehen!")
	os.Exit(0)
	return nil
}
