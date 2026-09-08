// Package bezugskalkulation contains formulars for calculations regarding the aquisition of wares
package bezugskalkulation

import "math"

// ZielEinkaufsPreis expects rabat to be the % float like 5 not the decimal float 0.05
func ZielEinkaufsPreis(listenpreis float64, rabat float64) float64 {
	if rabat == 0.0 {
		return listenpreis
	}

	rabatInDecimal := rabat / 100.0

	rabatInEuro := math.Round((listenpreis*rabatInDecimal)*100) / 100

	zieleinkaufspreis := math.Round((listenpreis-rabatInEuro)*100) / 100

	return zieleinkaufspreis
}

// BarEinkaufsPreis expects skonto to be the % flaot like 5 not the decimal float 0.05
func BarEinkaufsPreis(zieleinkaufspreis float64, skonto float64) float64 {
	if skonto == 0.0 {
		return zieleinkaufspreis
	}

	skontoInDecimal := skonto / 100.0

	skontoInEuro := math.Round((zieleinkaufspreis*skontoInDecimal)*100) / 100

	bareinkaufspreis := math.Round((zieleinkaufspreis-skontoInEuro)*100) / 100

	return bareinkaufspreis
}

func BezugsPreis(bareinkaufspreis float64, bezugskosten float64) float64 {
	return bareinkaufspreis + bezugskosten
}
