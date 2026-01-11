package config

// Config has unexported fields - should be fixed with -internal flag
type Config struct {
	host    string
	port    int
	Timeout int
}

// Service has mixed fields
type Service struct {
	Name    string
	enabled bool
}

func main() {
	// With -internal flag, these should be fixable
	cfg := Config{"localhost", 8080, 30}  // want "positional struct literal initialization is fragile"
	svc := Service{"api", true}           // want "positional struct literal initialization is fragile"

	_ = cfg
	_ = svc
}
