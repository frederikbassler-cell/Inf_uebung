package pointer

// Aufgabe 2: Swap zweier ints per Pointer
// Vertausche die Werte von a und b.
// Wenn einer der Pointer nil ist, soll nichts passieren.
func Swap(a, b *int) {
	if a == nil || b == nil{
		return  
	}
 s := 0 

 s = *a 

 *a = *b 

 *b = s 
 

}
