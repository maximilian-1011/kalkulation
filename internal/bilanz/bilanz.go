// Package bilanz contains the functions to calculate the Geschäftsbilanz
package bilanz

func CalculateBilanz(anlagevermögen, umflaufvermögen, forderungOffen, kassenbestand, bankvermögen, schuldenLang, schuldenKurz float64) float64 {
	privatVermögen := (anlagevermögen + umflaufvermögen + forderungOffen + kassenbestand + bankvermögen) - (schuldenLang + schuldenKurz)
	return privatVermögen
}
