// Package lagerkennzahlen contains the formulars to calculate the relevant values relating to storage
package lagerkennzahlen

import "fmt"

func Umschlagshäufigkeit(umsatz float64, druchschnittlicherLagerbestand float64) (float64, error) {
	if druchschnittlicherLagerbestand == 0.0 || umsatz == 0.0 {
		return 0.0, fmt.Errorf("zero division error")
	}

	return umsatz / druchschnittlicherLagerbestand, nil
}

func Lagerdauer(lagerumschalg float64) (float64, error) {
	if lagerumschalg == 0.0 {
		return 0.0, fmt.Errorf("zero division error")
	}

	return 360 / lagerumschalg, nil
}

func DruchschnittlicherLagerbestand(jahresanfangsbestand float64, monatsbestände []float64) (float64, error) {
	if len(monatsbestände) < 1 {
		return 0.0, fmt.Errorf("keine monatsbestände")
	}
	lagerbestand := jahresanfangsbestand

	for _, bestand := range monatsbestände {
		lagerbestand += bestand
	}

	if lagerbestand == 0.0 {
		return 0.0, fmt.Errorf("zero division error")
	}

	return lagerbestand / float64(len(monatsbestände)+1), nil
}

func Lagerzinssatz(jahreszins float64, lagerdauer float64) (float64, error) {
	if jahreszins == 0.0 || lagerdauer == 0.0 {
		return 0.0, fmt.Errorf("zero division error")
	}

	return (jahreszins * lagerdauer) / 360, nil
}

// Lagerzinsen expects lagerbestand in Euro
func Lagerzinsen(lagerbestand float64, lagerzinssatz float64) (float64, error) {
	if lagerbestand == 0.0 || lagerzinssatz == 0.0 {
		return 0.0, fmt.Errorf("zero division error")
	}

	return (lagerbestand * lagerzinssatz) / 100, nil
}
