package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/maximilain-1011/kalkulation/internal/bezugskalkulation"
)

func commandBezugskalkulation() error {
	parms := []string{"Listenpreis", "Rabatt", "Skonto", "Bezugskosten"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}
	zek := bezugskalkulation.ZielEinkaufsPreis(values["Listenpreis"], values["Rabatt"])
	bek := bezugskalkulation.BarEinkaufsPreis(zek, values["Skonto"])
	fmt.Println()
	fmt.Printf("Das Ergebnis ist: %.2f€\n", bezugskalkulation.BezugsPreis(bek, values["Bezugskosten"]))
	fmt.Println()
	return nil
}

func commandZieleinkaufspreis() error {
	parms := []string{"Listenpreis", "Rabatt"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Das Ergebins ist: %.2f€\n", bezugskalkulation.ZielEinkaufsPreis(values["Listenpreis"], values["Rabatt"]))
	fmt.Println()
	return nil
}

func commandBareinkaufspreis() error {
	parms := []string{"Listenpreis", "Rabatt", "Skonto"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Das Ergebins ist: %.2f€\n", bezugskalkulation.BarEinkaufsPreis(bezugskalkulation.ZielEinkaufsPreis(values["Listenpreis"], values["Rabatt"]), values["Skonto"]))
	fmt.Println()
	return nil
}

func commandSelbstkostenpreis() error {
	parms := []string{"Listenpreis", "Rabatt", "Skonto", "Bezugskosten", "Handlungskosten"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	zek := bezugskalkulation.ZielEinkaufsPreis(values["Listenpreis"], values["Rabatt"])
	bek := bezugskalkulation.BarEinkaufsPreis(zek, values["Skonto"])
	bk := bezugskalkulation.BezugsPreis(bek, values["Bezugskosten"])
	res := bezugskalkulation.SelbstkostenPreis(bk, values["Handlungskosten"])

	fmt.Println()
	fmt.Printf("Das Erbebnis ist: %.2f€\n", res)
	fmt.Println()
	return nil
}

func commandNettoVerkaufsPreis() error {
	parms := []string{"Listenpreis", "Rabatt", "Skonto", "Bezugskosten", "Handlungskosten", "Gewinnzuschlag"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	zek := bezugskalkulation.ZielEinkaufsPreis(values["Listenpreis"], values["Rabatt"])
	bek := bezugskalkulation.BarEinkaufsPreis(zek, values["Skonto"])
	bk := bezugskalkulation.BezugsPreis(bek, values["Bezugskosten"])
	skp := bezugskalkulation.SelbstkostenPreis(bk, values["Handlungskosten"])
	res := bezugskalkulation.NettoVerkaufsPreis(skp, values["Gewinnzuschlag"])

	fmt.Println()
	fmt.Printf("Das Erbebnis ist: %.2f€\n", res)
	fmt.Println()
	return nil
}

func commandBruttoVerkaufsPreis() error {
	parms := []string{"Listenpreis", "Rabatt", "Skonto", "Bezugskosten", "Handlungskosten", "Gewinnzuschlag", "Umsatzsteuer"}
	values, err := getValues(parms)
	if err != nil {
		return err
	}

	zek := bezugskalkulation.ZielEinkaufsPreis(values["Listenpreis"], values["Rabatt"])
	bek := bezugskalkulation.BarEinkaufsPreis(zek, values["Skonto"])
	bk := bezugskalkulation.BezugsPreis(bek, values["Bezugskosten"])
	skp := bezugskalkulation.SelbstkostenPreis(bk, values["Handlungskosten"])
	nvp := bezugskalkulation.NettoVerkaufsPreis(skp, values["Gewinnzuschlag"])
	res := bezugskalkulation.BruttoVerkaufsPreis(nvp, values["Umsatzsteuer"])

	fmt.Println()
	fmt.Printf("Das Erbebnis ist: %.2f€\n", res)
	fmt.Println()
	return nil
}

func getValues(parms []string) (map[string]float64, error) {
	scanner := bufio.NewScanner(os.Stdin)
	values := make(map[string]float64)
	for _, key := range parms {
		fmt.Printf("%s: ", key)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			return map[string]float64{}, err
		}

		words := cleanInput(scanner.Text())
		val, err := strconv.ParseFloat(words[0], 64)
		if err != nil {
			return map[string]float64{}, err
		}
		values[key] = val
	}

	return values, nil
}
