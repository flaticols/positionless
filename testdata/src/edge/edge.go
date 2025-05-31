package edge

import "fmt"

type Empty struct{}

type OneField struct {
	Value int
}

type MixedExport struct {
	Public  string
	private int
	Another string
}

type Embedded struct {
	fmt.Stringer
	Value int
}

type Anonymous struct {
	string
	int
}

func main() {
	// Should NOT trigger: empty struct
	e := Empty{}
	
	// Should trigger: single field struct
	o := OneField{42} // want "positional struct literal initialization is fragile"
	
	// Should NOT trigger: struct with unexported fields (can't fix)
	m := MixedExport{"public", 123, "another"}
	
	// Edge case: more values than fields (should NOT trigger - unsafe)
	// This would be a compile error anyway
	// bad := OneField{1, 2}
	
	// Should trigger: struct with embedded type
	emb := Embedded{nil, 100} // want "positional struct literal initialization is fragile"
	
	// Should NOT trigger: struct with anonymous fields (can't reference by name)
	anon := Anonymous{"text", 42}
	
	_ = e
	_ = o
	_ = m
	_ = emb
	_ = anon
}