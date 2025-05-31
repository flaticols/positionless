package nested

type Address struct {
	Street string
	City   string
	Zip    string
}

type Company struct {
	Name    string
	Address Address
	Founded int
}

type Employee struct {
	Name    string
	Company Company
	Salary  float64
}

func main() {
	// Should trigger: nested positional struct literals
	addr := Address{"123 Main St", "New York", "10001"} // want "positional struct literal initialization is fragile"
	
	// Should trigger: struct with nested struct field
	comp := Company{"TechCorp", addr, 2020} // want "positional struct literal initialization is fragile"
	
	// Should trigger: deeply nested
	emp := Employee{"John", comp, 75000.0} // want "positional struct literal initialization is fragile"
	
	// Mixed: outer positional, inner named (should trigger for outer only)
	comp2 := Company{"StartupInc", Address{ // want "positional struct literal initialization is fragile"
		Street: "456 Oak Ave",
		City:   "San Francisco",
		Zip:    "94102",
	}, 2021}
	
	// Should NOT trigger: all named
	emp2 := Employee{
		Name: "Jane",
		Company: Company{
			Name: "BigCorp",
			Address: Address{
				Street: "789 Pine Rd",
				City:   "Seattle",
				Zip:    "98101",
			},
			Founded: 2019,
		},
		Salary: 85000.0,
	}
	
	_ = addr
	_ = comp
	_ = emp
	_ = comp2
	_ = emp2
}