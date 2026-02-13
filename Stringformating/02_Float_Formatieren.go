package stringformating

import "fmt"

//Formatiere einen Preis mit zwei Nachkommastellen.
func FormatPrice(price float64) string {
	

	return fmt.Sprintf("Preis: %.2f €", price)
}
