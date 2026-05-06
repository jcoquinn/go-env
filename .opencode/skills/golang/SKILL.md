---
name: golang
description:
  Go coding conventions and idioms for this repository. Use when reading,
  writing, editing, or reviewing Go source (.go), go.mod, or go.sum files.
---

# Go Development Instructions

Follow idiomatic Go practices and community standards when writing Go code.
These instructions are based on [Effective Go](https://go.dev/doc/effective_go),
[Go Code Review Comments](https://go.dev/wiki/CodeReviewComments), and
[Google's Go Style Guide](https://google.github.io/styleguide/go/).

## Repository Context

Single package library: `env/env.go`. There is no `main`, no `cmd/`, no
`internal/`, no subpackages. All code lives in `package env`.

Go version: (see `.go-version` and `go.mod`).

## General Instructions

- Write simple, clear, and idiomatic Go code
- Favor clarity and simplicity over cleverness
- Follow the principle of least surprise
- Keep the happy path left-aligned (minimize indentation)
- Return early to reduce nesting; use `if condition { return }` to avoid else
  blocks
- Make the zero value useful
- Write self-documenting code with clear, descriptive names
- Document all exported types, functions, methods, and packages
- Leverage the Go standard library over custom implementations
- Avoid emoji in code, comments, and documentation

## Naming Conventions

### Packages

- Lowercase, single-word names; no underscores, hyphens, or mixedCaps
- Names describe what the package provides, not what it contains
- Avoid generic names like `util`, `common`, or `base`

### Package Declaration Rules (CRITICAL)

- **NEVER duplicate `package` declarations** — each Go file must have exactly
  one `package` line
- When editing an existing file: **preserve** the existing `package` declaration
- When creating a new file: check what package name other `.go` files in the
  same directory use and match it. In this repo that is always `env`.
- **NEVER** create files with multiple `package` lines

### Variables and Functions

- Use mixedCaps / MixedCaps (camelCase), not underscores
- Exported names start with a capital letter; unexported with lowercase
- Avoid stuttering (e.g. prefer `http.Server` over `http.HTTPServer`)
- Single-letter variables only for very short scopes (loop indices, etc.)

### Interfaces

- Use `-er` suffix for single-method interfaces (`Reader`, `Writer`)
- Keep interfaces small and focused (1–3 methods is ideal)
- Define interfaces near where they're used, not where they're implemented
- Don't export interfaces unless necessary

### Constants

- MixedCaps for exported, mixedCaps for unexported
- Group related constants in `const` blocks
- Use typed constants for better type safety

## Code Style and Formatting

- Format order (run on every changed `.go` file):
  1. `go tool goimports -w <file>`
  2. `go tool golines -w <file>` — wraps code lines at 100 chars; does **not**
     wrap comments
  3. `go tool gofumpt -w <file>`
- Wrap godoc and inline comments manually at **80 chars**; no formatter enforces
  this
- Add blank lines to separate logical groups of code
- Comments in complete English sentences; start with the name of the thing
  described
- Package comments: `// Package env …`
- Use `//` line comments for most comments; `/* */` only for package docs
- Document the _why_, not the _what_, unless the what is complex

## Error Handling

- Check errors immediately after the function call
- Don't ignore errors with `_` without a documented reason
- Wrap errors with context: `fmt.Errorf("doing X: %w", err)`
- Create custom error types when callers need to distinguish errors
- Error return is always the last return value
- Name error variables `err`
- Error messages: lowercase, no trailing punctuation
- Don't use panic for normal error handling
- Use `errors.Is` / `errors.As` for error inspection

## Type Safety and Language Features

- Define types to add meaning and type safety
- Prefer explicit type conversions
- Check both return values of type assertions
- Prefer specific types or generic type parameters with constraints over
  unconstrained types; use `any` (not `interface{}`) when an unconstrained type
  is required (Go 1.18+)

### Pointers vs Values

- Prefer pointer receivers; use value receivers when immutability is desired
- Be consistent within a type's method set
- Consider the zero value when choosing receiver type

### Interfaces and Composition

- Accept interfaces, return concrete types
- Use embedding for composition

## Concurrency

- Always know how a goroutine will exit
- Avoid goroutine leaks by ensuring cleanup
- Close channels from the sender side
- Use channels for communication, mutexes for protecting state
- Use `sync.Once` for one-time initialization

### WaitGroup (Go 1.26 — use `wg.Go`)

```go
var wg sync.WaitGroup
wg.Go(task1)
wg.Go(task2)
wg.Wait()
```

## Testing

- Test file uses `package env_test` (black-box / external package)
- Use `github.com/stretchr/testify` packages `assert` and `require` for test
  assertions; prefer `assert` by default; use `require` when subsequent
  assertions are meaningless if the current one fails (e.g. type assertion
  preconditions)
- Test naming: `Test<Type>_<Description>` (e.g. `TestInt_PanicsWhenUnset`)
- Run tests: `go test -v ./...`
- Use `t.Helper()` in test helper functions
- Clean up resources with `t.Cleanup()`
- Use `testing.TB` for helpers shared between tests and benchmarks

## Dependency Management

- Use `go mod tidy` to clean up unused dependencies
- Keep dependencies minimal
- Vendor only when necessary

## Security

- Validate all external input
- Use `crypto/rand` for random number generation
- Use TLS for network communication
- Don't implement custom cryptography

## Common Pitfalls to Avoid

- Not checking errors
- Ignoring race conditions
- Creating goroutine leaks
- Forgetting `defer` for cleanup
- Modifying maps concurrently
- Misunderstanding nil interfaces vs nil pointers
- Forgetting to close resources (files, connections)
- Using global variables unnecessarily
- **Creating duplicate `package` declarations** — always check existing files
  before adding a package declaration
