package main

import (
	"fmt"

	"github.com/maximilain-1011/kalkulation/internal/bilanz"
)

func commandBilanz() error {
	parms := []string{"Anlagevermögen", "Umlaufvermögen", "Offene Forderungen", "Kassenbestand", "Bankvermögen", "Langfristige Schulden", "Kurzfristige Schulden"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	fmt.Printf("Aus der Bilanz ergibt sich folgendes Privatvermögen: %.2f\n", bilanz.CalculateBilanz(values["Anlagevermögen"], values["Umlaufvermögen"], values["Offene Forderungen"], values["Kassenbestand"], values["Bankvermögen"], values["Langfristige Schulden"], values["Kurzfristige Schulden"]))
	return nil
}
