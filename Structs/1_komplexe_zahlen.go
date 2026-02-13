package structs

import "fmt"

// Struct zur Darstellung einer komplexen Zahl in Komponentenform
type ComplexNumber struct {
	Real float64
	Imag float64
}

// Gibt die komplexe Zahl als String zurück
func (c ComplexNumber) String() string {
	sign := "+"
	if c.Imag < 0 {
		sign = "-"
	}
    // WICHTIG: Hier abs() nutzen, sonst steht da z.B. "3 - -2i"
	return fmt.Sprintf("%.f %s %.fi", c.Real, sign, abs(c.Imag))
}

// Hilfsfunktion zur Rückgabe des Betrags eines float64
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// AddComplex addiert zwei komplexe Zahlen (a + bi) + (c + di)
func AddComplex(a, b ComplexNumber) ComplexNumber {
	var l ComplexNumber
	l.Real = a.Real + b.Real
	l.Imag = a.Imag + b.Imag
    // Kein Zusammenrechnen von Real und Imag hier!
	return l 
}

// MultiplyComplex multipliziert zwei komplexe Zahlen (a + bi) * (c + di)
// Formel: (ac - bd) + (ad + bc)i
func MultiplyComplex(a, b ComplexNumber) ComplexNumber {
	var l ComplexNumber
    
    // Realteil: ac - bd
	l.Real = (a.Real * b.Real) - (a.Imag * b.Imag)
    
    // Imaginärteil: ad + bc
	l.Imag = (a.Real * b.Imag) + (a.Imag * b.Real)
    
	return l
}