package stringformating

import "fmt"

// Schreibe eine Funktion, die Name und Alter zu einem formatierten String zusammensetzt.
func FormatPerson(name string, age int) string {

	return fmt.Sprintf("Name: %s, Alter: %d", name, age)
}
