package pointer

// Aufgabe 5: Slice per Funktion verändern
// Verdopple jeden Wert im Slice in-place.
// Wenn s == nil oder len(s)==0, soll nichts passieren.
func DoubleAll(s []int) {
	
	if len(s) == 0 {
		return 
	}

	for i := range s {

		s[i] *= 2 
	}
	
}
