package main

import (
	"fmt"

	"github.com/maximilain-1011/kalkulation/internal/lagerkennzahlen"
)

func commandUmschlagshäufigkeit() error {
	parms := []string{"Umsatz", "Durchschnittlicher Lagerbestand"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	res, err := lagerkennzahlen.Umschlagshäufigkeit(values["Umsatz"], values["Durchschnittlicher Lagerbestand"])
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Das Ergebnis ist: %.2f\n", res)
	fmt.Println()
	return nil
}

func commandLagerdauer() error {
	parms := []string{"Lagerumschlag"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	res, err := lagerkennzahlen.Lagerdauer(values["Lagerumschlag"])
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Das Ergebnis ist: %.2f\n", res)
	fmt.Println()
	return nil
}

func commandDurchschnittlicherLagerbestand() error {
	parms := []string{"Jahresanfangsbestand", "Jan", "Feb", "März", "April", "Mai", "Juni", "Juli", "Aug", "Sep", "Okt", "Nov", "Dec"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	monatsbestände := []float64{}
	for key, val := range values {
		if key == "Jahresanfangsbestand" {
			continue
		}

		monatsbestände = append(monatsbestände, val)
	}

	res, err := lagerkennzahlen.DruchschnittlicherLagerbestand(values["Jahresanfangsbestand"], monatsbestände)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Das Ergebnis ist: %.2f", res)
	fmt.Println()
	return nil
}

func commandLagerzinssatz() error {
	parms := []string{"Jahreszins", "Lagerdauer"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	res, err := lagerkennzahlen.Lagerzinssatz(values["Jahreszins"], values["Lagerdauer"])
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Das Ergebnis ist: %.2f\n", res)
	fmt.Println()
	return nil
}

func commandLagerzins() error {
	parms := []string{"Lagerbestand", "Lagerzinssatz"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	res, err := lagerkennzahlen.Lagerzinsen(values["Lagerbestand"], values["Lagerzinssatz"])
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Das Ergebnis ist: %.2f\n", res)
	fmt.Println()
	return nil
}
