package lagerkennzahlen

import "fmt"

func Umschlagshäufigkeit(umsatz float64, druchschnittlicher_lagerbestand float64) (float64, error) {
	if druchschnittlicher_lagerbestand == 0.0 || umsatz == 0.0 {
		return 0.0, fmt.Errorf("Zero division error\n")
	}

	return umsatz/druchschnittlicher_lagerbestand, nil
}

func Lagerdauer(lagerumschalg float64) (float64, error) {
	if lagerumschalg == 0.0 {
		return 0.0, fmt.Errorf("Zero division error\n")
	}

	return 360/lagerumschalg, nil
}

func DruchschnittlicherLagerbestand(jahresanfangsbestand float64, monatsbestände []float64) (float64, error) {
	if len(monatsbestände) < 1 {
		return 0.0, fmt.Errorf("Keine monatsbestände\n")
	}
	lagerbestand := jahresanfangsbestand

	for _, bestand := range monatsbestände {
		lagerbestand += jahresanfangsbestand
	}

	if lagerbestand == 0.0 {
		return 0.0, fmt.Errorf("Zero division error\n")
	}

	return lagerbestand/len(monatsbestände)+1, nil
}

func Lagerzinssatz(jahreszins float64, lagerdauer float64) (float64, error) {
	if jahreszins == 0.0 || lagerdauer == 0.0 {
		return 0.0, fmt.Errorf("Zero division error\n")
	}

	return (jahreszins*lagerdauer)/360, nil
}

// lagerbestand in Euro
func Lagerzinsen(lagerbestand float64, lagerzinssatz float64) (float64, error) {
	if lagerbestand == 0.0 || lagerzinssatz == 0.0 {
		return 0.0, fmt.Errorf("Zero division error\n")
	}

	return (lagerbestand*lagerzinssatz)/100, nil
}
