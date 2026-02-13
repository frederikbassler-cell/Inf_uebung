package slices

// RemoveShortWords entfernt alle Wörter aus dem Slice,
// die kürzer sind als die angegebene Mindestlänge minLen.
// Das Slice wird direkt über einen Pointer verändert.
func RemoveShortWords(words []string, minLen int) []string {
var new []string

	for _,v := range words {
		if  minLen <= len(v) {
			new = append(new, v)
			
		}
	}





	return new
}
