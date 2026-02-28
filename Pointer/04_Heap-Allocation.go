package pointer

// Aufgabe 4: new / Heap-Allocation
// Erstelle und returne eine *Person mit den übergebenen Werten.
// (Nutze new(Person) oder &Person{...})
func NewPerson(name string, alter int) *Person {
	p := new(Person) // p ist jetzt schon ein Pointer (*Person)
	
	p.Name = name    // Wir befüllen das leere Struct
	p.Alter = alter
	
	return p

}

