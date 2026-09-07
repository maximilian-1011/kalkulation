package main

import (
	"fmt"

	"github.com/maximilain-1011/kalkulation/internal/bezugskalkulation"
)

func commandBezugskalkulation() error {
	fmt.Printf("Das ergebnis ist: %v\n", bezugskalkulation.BezugsPreis(10.0, 5.0))
	return nil
}

func commandZieleinkaufspreis() error {
	return nil
}

func commandBareinkaufspreis() error {
	return nil
}
