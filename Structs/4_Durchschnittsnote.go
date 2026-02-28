package structs

// Struct zur Darstellung einer Bewertung
type Rating struct {
	Name  string // Name der bewertenden Person
	Score int    // Bewertung zwischen z. B. 1 und 5
}

// AverageScore berechnet den Durchschnittswert aller Bewertungen
func AverageScore(ratings []Rating) float64 {
	points := 0
	count:= 0 
	var average float64

for _,v := range ratings{

	if v.Score >=1 && v.Score <= 5 {
		count++
		points += v.Score}
	
}


average = float64(points / count)

	return average // Platzhalter
}
