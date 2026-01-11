package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// Test package names
const (
	pkgBasic    = "basic"
	pkgNested   = "nested"
	pkgPointer  = "pointer"
	pkgEdge     = "edge"
	pkgGenerated = "generated"
	pkgUnexported = "unexported"
	pkgInternal   = "internal/config"
	pkgIgnore     = "ignore"
	pkgJSONTest   = "jsontest"
)

// Test ignore patterns
const testIgnorePatterns = "MixedExport,Anonymous"

func TestPositionless(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, pkgBasic, pkgNested, pkgPointer, pkgEdge)
}

func TestPositionlessWithGenerated(t *testing.T) {
	includeGenerated = true
	defer func() { includeGenerated = false }()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, pkgGenerated)
}

func TestPositionlessWithUnexported(t *testing.T) {
	includeUnexported = true
	defer func() { includeUnexported = false }()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, pkgUnexported)
}

func TestPositionlessWithInternal(t *testing.T) {
	detectInternal = true
	defer func() { detectInternal = false }()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, pkgInternal)
}

func TestPositionlessWithIgnore(t *testing.T) {
	ignorePatterns = testIgnorePatterns
	defer func() { ignorePatterns = "" }()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, pkgIgnore)
}

func TestPositionlessJSONOutput(t *testing.T) {
	outputFormat = outputJSON
	defer func() { outputFormat = outputText }()

	testdata := analysistest.TestData()
	// JSON mode outputs to stderr and skips pass.Report()
	// Use jsontest package which has no // want annotations
	analysistest.Run(t, testdata, Analyzer, pkgJSONTest)
}