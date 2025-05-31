# positionless

A Go static analyzer that detects positional struct literal initialization and suggests converting them to named field initialization for better code maintainability.

## Why?

Positional struct literals are fragile and can lead to bugs when struct fields are reordered or new fields are added. This analyzer helps you find and fix these issues automatically.

```go
// Bad - positional initialization
person := Person{"John", 30, "john@example.com"}

// Good - named field initialization
person := Person{
    Name:  "John",
    Age:   30,
    Email: "john@example.com",
}
```

## Installation

```bash
go install github.com/flaticols/positionless@latest
```

## Usage

### As a standalone tool

```bash
# Analyze current directory
positionless ./...

# Analyze specific package
positionless ./pkg/mypackage

# Include generated files (excluded by default)
positionless -generated ./...

# Apply suggested fixes automatically
positionless -fix ./...
```

### Using with other tools

This tool pairs well with `fieldalignment` analyzer. Run `positionless` first to convert positional literals to named fields, then run `fieldalignment` to optimize struct memory layout:

```bash
# First, fix positional initialization
positionless -fix ./...

# Then, optimize field alignment
fieldalignment -fix ./...
```

### In your editor

Most Go editors support running custom analyzers. Configure your editor to run this analyzer for real-time feedback.

## How it works

The analyzer:

1. Scans your Go code for struct literal initialization
2. Identifies positional initialization patterns
3. Suggests fixes that convert to named field initialization
4. Can automatically apply fixes with the `-fix` flag
5. Preserves your original values and formatting
6. Only processes exported fields (respects Go's visibility rules)
7. Skips generated files by default (use `-generated` to include them)

## Example

Given this code:

```go
type Config struct {
    Host     string
    Port     int
    Timeout  time.Duration
    RetryMax int
}

cfg := Config{"localhost", 8080, 5 * time.Second, 3}
```

The analyzer will suggest:

```go
cfg := Config{
    Host:     "localhost",
    Port:     8080,
    Timeout:  5 * time.Second,
    RetryMax: 3,
}
```

## License

MIT License - see LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
