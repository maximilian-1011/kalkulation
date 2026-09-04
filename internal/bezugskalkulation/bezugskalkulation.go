package bezugskalkulation

// rabat in % like 5% will be converted into 0.05
func ZielEinkaufsPreis(listenpreis float64, rabat float64) float64 {
	if rabat == 0.0 {
		return listenpreis
	}

	rabatInDecimal := rabat / 100.0

	rabatInEuro := listenpreis * rabatInDecimal

	zieleinkaufspreis := listenpreis - rabatInEuro

	return zieleinkaufspreis
}

// skonto in % like 5% will be converted into decimal 0.05
func BarEinkaufsPreis(zieleinkaufspreis float64, skonto float64) float64 {
	if skonto == 0.0 {
		return zieleinkaufspreis
	}

	skontoInDecimal := skonto / 100.0

	skontoInEuro := zieleinkaufspreis * skontoInDecimal

	bareinkaufspreis := zieleinkaufspreis - skontoInEuro

	return bareinkaufspreis
}

func BezugsPreis(bareinkaufspreis float64, bezugskosten float64) float64 {
	return bareinkaufspreis + bezugskosten
}
