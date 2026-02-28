package pointer

// Aufgabe 7: Pointer auf Pointer
// Setze den int-Wert, auf den **pp letztlich zeigt, auf v.
// Wenn pp == nil oder *pp == nil, soll nichts passieren.
func SetThroughDoublePointer(pp **int, v int) {

	if pp == nil || *pp == nil{
		return
	}

	**pp = v 
	

}


