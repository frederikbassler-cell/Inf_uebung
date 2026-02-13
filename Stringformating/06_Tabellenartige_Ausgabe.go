package stringformating

import (
	"fmt"
	"strings"
)

// Formatiere Name und Punktzahl spaltenbündig.
func FormatScore(name string, score int) string {
	länge := len(name)
	if länge > 10 {
		return fmt.Sprintf("%s %d", name, score)
	}
	 
	spaces := 10 - länge
	

	ein := strings.Repeat(" ", spaces)
	return fmt.Sprintf("%s%s %d", name, ein, score)
}
