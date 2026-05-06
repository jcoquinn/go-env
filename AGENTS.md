# AGENTS.md

## Commands

```sh
# Unit tests
go test -v ./...

# Format (run all three on changed files)
go tool goimports -w <file>
go tool golines -w <file>
go tool gofumpt -w <file>

# Lint
go tool staticcheck ./...

# All pre-commit hooks
pre-commit run --all-files

# Install declared Go tools
go install tool
```

CI only runs `pre-commit run --all-files` — there is no separate `go test` CI
job.

## Structure

Single package library: `env/env.go`. No `main`, no `cmd/`, no subpackages.

## Conventions

- Test file uses `package env_test` (black-box)
- Use `github.com/stretchr/testify` packages `assert` and `require`; prefer
  `assert` by default; use `require` when subsequent assertions are meaningless
  if the current one fails (e.g. type assertion preconditions)
- Test naming: `Test<Type>_<Description>` (e.g. `TestInt_PanicsWhenUnset`)
- Formatters run in order: `goimports` → `golines` → `gofumpt`
- YAML files must start with `---` (`yamlfmt` `include_document_start: true`)
- Markdown prose wrap at 80 chars (`prettier`)
- Go version: 1.26 (see `.go-version`)
- Use the `golang` skill (`.opencode/skills/golang/SKILL.md`) when reading,
  writing, or reviewing `.go`, `go.mod`, or `go.sum` files
