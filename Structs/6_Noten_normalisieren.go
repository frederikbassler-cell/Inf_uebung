package structs

// Student repräsentiert eine*n Schüler*in mit Name und Note.
type Student struct {
	Name  string
	Grade int // Note von 0 bis 100
}

// NormalizeGrades überprüft die Noten der Schüler*innen
// und passt sie mithilfe von Pointern an:
// - Noten über 100 werden auf 100 gesetzt
// - Noten unter 0 werden auf 0 gesetzt
func NormalizeGrades(students []Student) {

	for i := range students {

		p := &students[i].Grade
		if *p > 100 {
			*p = 100
		}

		if *p < 0 {
			*p = 0
		}


	}
}
