// Package bezugskalkulation contains formulars for calculations regarding the aquisition of wares
package bezugskalkulation

import "math"

// ZielEinkaufsPreis expects rabatt to be the % float like 5 not the decimal float 0.05
func ZielEinkaufsPreis(listenpreis, rabatt float64) float64 {
	if rabatt == 0.0 {
		return listenpreis
	}

	rabattInDecimal := rabatt / 100.0

	rabattInEuro := math.Round((listenpreis*rabattInDecimal)*100) / 100

	zieleinkaufspreis := math.Round((listenpreis-rabattInEuro)*100) / 100

	return zieleinkaufspreis
}

// BarEinkaufsPreis expects skonto to be the % flaot like 5 not the decimal float 0.05
func BarEinkaufsPreis(zieleinkaufspreis, skonto float64) float64 {
	if skonto == 0.0 {
		return zieleinkaufspreis
	}

	skontoInDecimal := skonto / 100.0

	skontoInEuro := math.Round((zieleinkaufspreis*skontoInDecimal)*100) / 100

	bareinkaufspreis := math.Round((zieleinkaufspreis-skontoInEuro)*100) / 100

	return bareinkaufspreis
}

func BezugsPreis(bareinkaufspreis, bezugskosten float64) float64 {
	return bareinkaufspreis + bezugskosten
}

func SelbstkostenPreis(bezugsPreis, handlungsKosten float64) float64 {
	handlungsKostenProzent := (bezugsPreis / 100) * handlungsKosten

	handlungsKostenEuro := math.Round(handlungsKostenProzent*100) / 100

	selbstKostenPreis := bezugsPreis + handlungsKostenEuro

	return selbstKostenPreis
}

func NettoVerkaufsPreis(selbstKostenPreis, gewinnZuschlag float64) float64 {
	gewinnZuschlagDecimal := gewinnZuschlag / 100.0

	gewinnZuschlagEuro := math.Round((selbstKostenPreis*gewinnZuschlagDecimal)*100) / 100

	nettoVerkaufsPreis := selbstKostenPreis + gewinnZuschlagEuro

	return nettoVerkaufsPreis
}

func BruttoVerkaufsPreis(nettoVerkaufsPreis, umsatzsteuer float64) float64 {
	umsatzsteuerDecimal := umsatzsteuer / 100

	umsatzsteuerInEuro := math.Round((nettoVerkaufsPreis*umsatzsteuerDecimal)*100) / 100

	bruttoVerkaufsPreis := nettoVerkaufsPreis + umsatzsteuerInEuro

	return bruttoVerkaufsPreis
}
