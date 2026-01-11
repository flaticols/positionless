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

func TestPositionlessWithUnexported(t *testing.T) {
	includeUnexported = true
	defer func() { includeUnexported = false }()

	testdata := analysistest.TestData()
	// With -unexported, MixedExport and Anonymous should have fixes
	analysistest.Run(t, testdata, Analyzer, "unexported")
}

func TestPositionlessWithInternal(t *testing.T) {
	detectInternal = true
	defer func() { detectInternal = false }()

	testdata := analysistest.TestData()
	// With -internal, internal/ packages should have fixes for unexported fields
	analysistest.Run(t, testdata, Analyzer, "internal/config")
}

func TestPositionlessWithIgnore(t *testing.T) {
	ignorePatterns = "MixedExport,Anonymous"
	defer func() { ignorePatterns = "" }()

	testdata := analysistest.TestData()
	// With ignore patterns, MixedExport and Anonymous should be skipped
	analysistest.Run(t, testdata, Analyzer, "ignore")
}