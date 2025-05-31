package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPositionless(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "basic", "nested", "pointer", "edge")
}

func TestPositionlessWithGenerated(t *testing.T) {
	includeGenerated = true
	defer func() { includeGenerated = false }()
	
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "generated")
}