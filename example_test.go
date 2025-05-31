package main_test

import (
	"fmt"
)

func Example() {
	// This example shows how to use the positionless analyzer
	// In practice, you would typically run it via the command line
	
	// First install the analyzer:
	// go install github.com/flaticols/positionless@latest
	
	// Then run it on your code:
	fmt.Println("Run: positionless ./...")
	
	// Output:
	// Run: positionless ./...
}

func Example_withFix() {
	// Example showing how to apply fixes automatically
	fmt.Println("Apply fixes: positionless -fix ./...")
	
	// Output:
	// Apply fixes: positionless -fix ./...
}

func Example_withGoVet() {
	// Example showing usage with go vet
	fmt.Println("With go vet: go vet -vettool=$(which positionless) ./...")
	
	// Output:
	// With go vet: go vet -vettool=$(which positionless) ./...
}