package jsontest

// Simple struct for JSON output testing.
type Person struct {
	Name string
	Age  int
}

func example() {
	// This triggers in JSON mode - no diagnostic expected since JSON skips pass.Report
	p := Person{"John", 30}
	_ = p
}
