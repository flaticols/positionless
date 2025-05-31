package pointer

type Node struct {
	Value int
	Next  *Node
}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func main() {
	// Should trigger: pointer to struct with positional init
	n1 := &Node{1, nil} // want "positional struct literal initialization is fragile"
	
	// Should trigger: struct with pointer field
	n2 := Node{2, n1} // want "positional struct literal initialization is fragile"
	
	// Should trigger: recursive structure (will match multiple times for nested literals)
	tree := &Tree{10, &Tree{5, nil, nil}, &Tree{15, nil, nil}} // want "positional struct literal initialization is fragile" "positional struct literal initialization is fragile" "positional struct literal initialization is fragile"
	
	// Should NOT trigger: named fields with pointers
	n3 := &Node{
		Value: 3,
		Next:  n2.Next,
	}
	
	// Should NOT trigger: named fields in recursive structure
	tree2 := &Tree{
		Value: 20,
		Left: &Tree{
			Value: 10,
			Left:  nil,
			Right: nil,
		},
		Right: &Tree{
			Value: 30,
			Left:  nil,
			Right: nil,
		},
	}
	
	_ = n1
	_ = n2
	_ = n3
	_ = tree
	_ = tree2
}