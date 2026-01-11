# positionless

A Go static analyzer that detects positional struct literal initialization and suggests converting them to named field initialization for better code maintainability.

## Why?

Positional struct literals are fragile and can lead to bugs when struct fields are reordered or new fields are added. This analyzer helps you find and fix these issues automatically.

> [!WARNING]
> Positional struct literals break when fields are reordered or new fields are added in the middle of a struct

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

### Go

```bash
go install github.com/flaticols/positionless@latest
```

### Homebrew

```bash
brew install flaticols/apps/positionless
```

## Usage

### As a standalone tool

```bash
# Analyze current directory
positionless ./...

# Analyze specific package
positionless ./pkg/mypackage

# Apply suggested fixes automatically
positionless -fix ./...

# Include generated files (excluded by default)
positionless -generated ./...

# Fix structs with unexported fields in internal packages
positionless -fix -internal ./...

# Skip specific struct types
positionless -ignore="ConfigTest,*Mock" ./...

# JSON output for tooling integration
positionless -output=json ./... 2>&1
```

### With go vet

You can use this analyzer with `go vet`:

```bash
# Run the analyzer with go vet
go vet -vettool=$(which positionless) ./...

# Apply fixes with go vet
go vet -vettool=$(which positionless) -fix ./...
```

> [!TIP]
> Add this to your CI/CD pipeline to catch positional struct literals early

### Using with other tools

This tool pairs well with `fieldalignment` analyzer. Run `positionless` first to convert positional literals to named fields, then run `fieldalignment` to optimize struct memory layout:

```bash
# First, fix positional initialization
positionless -fix ./...

# Then, optimize field alignment
fieldalignment -fix ./...
```

> [!NOTE]
> Running `positionless` before `fieldalignment` ensures that field reordering won't break your code

### In your editor

Most Go editors support running custom analyzers. Configure your editor to run this analyzer for real-time feedback.

### With golangci-lint v2

`positionless` supports [golangci-lint v2 module plugins](https://golangci-lint.run/docs/plugins/module-plugins/).

**Step 1:** Create `.custom-gcl.yml` in your project root:

```yaml
version: v2.1.2
plugins:
  - module: 'github.com/flaticols/positionless'
    import: 'github.com/flaticols/positionless'
    version: v2.0.0
```

**Step 2:** Build custom golangci-lint:

```bash
# Build custom binary with positionless (requires golangci-lint v2 installed)
golangci-lint custom
```

**Step 3:** Configure `.golangci.yml`:

```yaml
version: "2"
linters:
  enable:
    - positionless
  settings:
    custom:
      positionless:
        type: "module"
        description: Detect positional struct literals
        # Pass flags via settings (optional)
        settings:
          generated: false
          unexported: false
          internal: true
          ignore: ""
          output: "text"
```

**Step 4:** Run:

```bash
./custom-gcl run ./...
```

> [!NOTE]
> Module plugins are the recommended approach for golangci-lint v2. No toolchain version matching required.

### As a GitHub Action

You can use `positionless` in your GitHub workflows to automatically check for positional struct literals:

```yaml
name: Code Analysis
on: [push, pull_request]

jobs:
  analyze:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      # Check for positional struct literals
      - uses: flaticols/positionless@v2
```

To automatically fix issues and commit the changes:

```yaml
name: Auto-fix Positional Literals
on: [push]

jobs:
  fix:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      # Fix positional struct literals
      - uses: flaticols/positionless@v2
        with:
          fix: true
          
      # Commit changes if any
      - uses: stefanzweifel/git-auto-commit-action@v5
        with:
          commit_message: 'fix: convert positional struct literals to named fields'
```

#### Action Inputs

| Input | Description | Default |
|-------|-------------|---------|
| `path` | Path to analyze | `./...` |
| `fix` | Apply suggested fixes automatically | `false` |
| `include-generated` | Include generated files in analysis | `false` |
| `include-unexported` | Include structs with unexported fields in fixes | `false` |
| `include-internal` | Auto-allow unexported fields in `internal/` packages | `false` |
| `ignore` | Comma-separated patterns to skip (e.g., `ConfigTest,*Mock`) | `` |
| `output` | Output format: `text` or `json` | `text` |
| `version` | Version of positionless to use | `latest` |

#### Action Outputs

| Output | Description |
|--------|-------------|
| `findings-count` | Number of positional struct literals found |
| `fixed-count` | Number of fixes applied (when fix is enabled) |
| `exit-code` | Exit code from the analyzer (0=success, 3=findings) |
| `version` | Version of positionless used |

#### Example Output

Here's what the GitHub Action output looks like when it detects positional struct literals:

```
Run flaticols/positionless@v2
Run # Determine version
Fetching latest version...
Downloading positionless v2 for Linux_x86_64...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0

100 2669k  100 2669k    0     0  7746k      0 --:--:-- --:--:-- --:--:-- 7746k
Run FLAGS=""
Running: positionless  ./...
Error: /home/runner/work/bump/bump/semver/version_test.go:19:18: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:20:23: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:21:24: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:22:29: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:23:35: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:36:19: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:37:19: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:38:24: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:39:26: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:40:25: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:41:30: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:42:36: positional struct literal initialization is fragile
Error: /home/runner/work/bump/bump/semver/version_test.go:43:18: positional struct literal initialization is fragile
Error: Process completed with exit code 3.
```

When issues are found, the action will fail with exit code 3, causing your CI pipeline to fail. This helps catch positional struct literals before they're merged into your main branch.

#### Real-world Example

Here's an actual fix that `positionless` would apply to the code from the example above:

```diff
func TestParse(t *testing.T) {
        tests := []parseTestsInput{
-               {nil, "1.2.3", Version{nil, nil, 1, 2, 3}, false},
-               {nil, "1.2.3-beta", Version{[]string{"beta"}, nil, 1, 2, 3}, false},
-               {nil, "1.2.3+build", Version{nil, []string{"build"}, 1, 2, 3}, false},
-               {nil, "1.2.3-beta+build", Version{[]string{"beta"}, []string{"build"}, 1, 2, 3}, false},
-               {nil, "1.2.3-beta.1+build.123", Version{[]string{"beta", "1"}, []string{"build", "123"}, 1, 2, 3}, false},
+               {nil, "1.2.3", Version{
+                       Prerelease: nil,
+                       Metadata:   nil,
+                       Major:      1,
+                       Minor:      2,
+                       Patch:      3,
+               }, false},
+               {nil, "1.2.3-beta", Version{
+                       Prerelease: []string{"beta"},
+                       Metadata:   nil,
+                       Major:      1,
+                       Minor:      2,
+                       Patch:      3,
+               }, false},
+               {nil, "1.2.3+build", Version{
+                       Prerelease: nil,
+                       Metadata:   []string{"build"},
+                       Major:      1,
+                       Minor:      2,
+                       Patch:      3,
+               }, false},
+               {nil, "1.2.3-beta+build", Version{
+                       Prerelease: []string{"beta"},
+                       Metadata:   []string{"build"},
+                       Major:      1,
+                       Minor:      2,
+                       Patch:      3,
+               }, false},
+               {nil, "1.2.3-beta.1+build.123", Version{
+                       Prerelease: []string{"beta", "1"},
+                       Metadata:   []string{"build", "123"},
+                       Major:      1,
+                       Minor:      2,
+                       Patch:      3,
+               }, false},
                {ErrMalformedCore, "78", Version{}, true},
                {ErrMalformedCore, "1.2", Version{}, true},
                {ErrMalformedCore, "1.2.3.4", Version{}, true},
```

Running `positionless -fix ./...` would automatically apply these changes, making your code more maintainable and resistant to struct field reordering.

## How it works

The analyzer:

1. Scans your Go code for struct literal initialization
2. Identifies positional initialization patterns
3. Reports all positional struct literals as warnings
4. Suggests fixes that convert to named field initialization
5. Can automatically apply fixes with the `-fix` flag
6. Preserves your original values and formatting
7. Skips generated files by default (use `-generated` to include them)

### Flags

| Flag | Description |
|------|-------------|
| `-fix` | Apply suggested fixes automatically |
| `-generated` | Include generated files in analysis |
| `-unexported` | Include structs with unexported fields in fixes |
| `-internal` | Auto-allow unexported fields in `internal/` packages |
| `-ignore=PATTERN` | Skip structs matching pattern (comma-separated, supports globs) |
| `-output=FORMAT` | Output format: `text` (default) or `json` (golangci-lint compatible) |

> [!TIP]
> Use `-internal` when analyzing your own internal packages where you control all the code

### Exit codes

| Code | Meaning |
|------|---------|
| 0 | No issues found |
| 1 | Error occurred |
| 3 | Issues found (useful for CI pipelines) |

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

> [!IMPORTANT]
> By default, the analyzer reports all positional structs but only auto-fixes those with exported fields. Use `-unexported` or `-internal` to enable fixes for structs with unexported fields.

## License

MIT License - see LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
