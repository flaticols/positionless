package basic

import "time"

type Person struct {
	Name  string
	Age   int
	Email string
}

type Config struct {
	Host     string
	Port     int
	Timeout  time.Duration
	Enabled  bool
	RetryMax int
}

func main() {
	// Should trigger: positional struct literal
	p1 := Person{"John", 30, "john@example.com"} // want "positional struct literal initialization is fragile"
	
	// Should trigger: positional with pointer
	p2 := &Person{"Jane", 25, "jane@example.com"} // want "positional struct literal initialization is fragile"
	
	// Should NOT trigger: already using named fields
	p3 := Person{
		Name:  "Bob",
		Age:   35,
		Email: "bob@example.com",
	}
	
	// Should NOT trigger: partial named fields
	p4 := Person{Name: "Alice", Age: 28}
	
	// Should trigger: complex struct
	cfg := Config{"localhost", 8080, 5 * time.Second, true, 3} // want "positional struct literal initialization is fragile"
	
	// Should NOT trigger: empty struct literal
	p5 := Person{}
	
	_ = p1
	_ = p2
	_ = p3
	_ = p4
	_ = p5
	_ = cfg
}