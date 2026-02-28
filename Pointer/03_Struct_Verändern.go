package pointer

// Aufgabe 3: Struct per Pointer verändern
type Person struct {
	Name  string
	Alter int
}

// Erhöhe das Alter der Person um 1.
// Wenn p == nil, soll nichts passieren.
func Birthday(p *Person) {

	if p == nil {
		return
	}

	(*p).Alter += 1 
 
}




