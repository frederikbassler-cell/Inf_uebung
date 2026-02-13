package stringformating

import "fmt"

//Gib den Wert und den Typ einer beliebigen Variable aus.
func DescribeValue(v interface{}) string {

	return fmt.Sprintf("Wert: %v | Typ: %T", v, v)
}
