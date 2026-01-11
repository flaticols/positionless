package unexported

// MixedExport has both exported and unexported fields
type MixedExport struct {
	Public  string
	private int
	Another string
}

// AllUnexported has only unexported fields
type AllUnexported struct {
	name string
	age  int
}

func main() {
	// With -unexported flag, these should be fixable
	m := MixedExport{"public", 123, "another"} // want "positional struct literal initialization is fragile"
	a := AllUnexported{"john", 30}             // want "positional struct literal initialization is fragile"

	_ = m
	_ = a
}
