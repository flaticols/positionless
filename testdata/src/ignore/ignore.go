package ignore

// MixedExport - should be ignored when pattern matches
type MixedExport struct {
	Public  string
	private int
}

// Anonymous - should be ignored when pattern matches
type Anonymous struct {
	string
	int
}

// Person - should NOT be ignored
type Person struct {
	Name string
	Age  int
}

func main() {
	// MixedExport and Anonymous should be ignored due to -ignore flag
	m := MixedExport{"public", 123}
	anon := Anonymous{"text", 42}

	// Person should still trigger
	p := Person{"John", 30} // want "positional struct literal initialization is fragile"

	_ = m
	_ = anon
	_ = p
}
