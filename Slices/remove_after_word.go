package slices

// CutAfterWord kürzt das Slice so, dass es nur die Wörter
// bis einschließlich des ersten Vorkommens von `target` enthält.
// Falls `target` nicht vorkommt, bleibt das Slice unverändert.
// Das Slice wird über einen Pointer übergeben und direkt verändert.
func CutAfterWord(words []string, target string) []string {
	var new []string

	for _,v := range words{
		if v == target {
			new = append(new, v)
			return new
		}else{
		new = append(new, v)
		}
		
	}

	return words
}
