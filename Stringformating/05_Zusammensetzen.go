package stringformating

import "fmt"

//Formatiere folgenden Text ohne Stringverkettung (+):
func UserText(name string, age int) string {
	
	return fmt.Sprintf("Der Benutzer %s ist %d Jahre alt.", name, age)
}
