package pointer

// Aufgabe 6: Nil-Pointer sicher dereferenzieren
// Wenn p == nil, returne (0, false).
// Sonst returne (*p, true).
func SafeDeref(p *int) (int, bool) {
	
	if p == nil {
		return 0, false 
	}else{
		return *p,true
	}

}
