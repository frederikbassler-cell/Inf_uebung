package pointer

// Aufgabe 1: Verdoppeln per Pointer
// Schreibe eine Funktion, die den Wert, auf den p zeigt, verdoppelt.
// Wenn p == nil, soll nichts passieren (keine Panic).
func DoubleViaPointer(p *int) {
	if p == nil {
		return  
	}

	*p *= 2	
}
