package structs

func ExampleAppendNode() {
	var head *Node = nil

	AppendNode(&head, "Go")
	AppendNode(&head, "ist")
	AppendNode(&head, "toll")
	AppendNode(&head, "Lüge")

	PrintList(head)

	// Output:
	// Go
	// ist
	// toll
	// Lüge
}
