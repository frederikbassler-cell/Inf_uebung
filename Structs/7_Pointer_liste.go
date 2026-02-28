package structs

import "fmt"

// Node repräsentiert ein Element einer einfach verketteten Liste.
type Node struct {
	Value string
	Next  *Node
}

// AppendNode fügt einen neuen Knoten mit dem gegebenen Wert
// ans Ende der Liste an. headPtr ist ein Pointer auf das
// erste Listenelement (also ein Pointer auf einen Pointer).
// Wenn die Liste leer ist (head == nil), wird der neue Knoten
// zum Kopf der Liste.
func AppendNode(headPtr **Node, value string) {
	newNode := &Node{Value: value, Next: nil}

	if *headPtr == nil {
		*headPtr = newNode
		return
	}

	current := *headPtr
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// PrintList gibt alle Werte der Liste zeilenweise aus.
func PrintList(n *Node) {
	for n != nil {
		fmt.Println(n.Value) // Hier passiert die Ausgabe für den Test
		n = n.Next
	}
}

