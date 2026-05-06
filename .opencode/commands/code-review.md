---
description: Review staged changes for bugs, gaps, and convention drift
agent: plan
---

Use the `golang` skill before reviewing.

Perform a thorough code review of the staged changes below (or files matched by
$ARGUMENTS if provided). Use codebase search for additional context as needed.

Organize findings by the sections below. For each finding, include:

- The file and line reference
- What the issue is and why it matters
- A concrete suggestion or corrected snippet

Only report genuine issues — avoid nitpicks that don't affect correctness,
clarity, or maintainability.

Project conventions reference: @AGENTS.md

Staged diff: !`git diff --cached $ARGUMENTS`

Working tree status for context: !`git status --short`

---

## Go Code Review

Focus on `*.go`. Apply the `golang` skill conventions.

**Correctness**

- Logic errors, off-by-one errors, incorrect assumptions
- Improper error handling: errors silently ignored, errors logged AND returned
  (pick one), wrong sentinel error (`ErrParse` vs `ErrNotExist`)
- Panics are intentional for unset/unparseable env vars in this library — flag
  only misuse outside that pattern
- Incorrect use of goroutines, channels, or shared state

**Clarity and Idioms**

- Non-idiomatic patterns; prefer Effective Go / Go Code Review Comments
- Missing or misleading doc comments on exported types and functions
- Names that don't clearly express intent
- Use of `interface{}` instead of `any`

**Testing**

- Missing tests for new public functions
- Tests not following `Test<Type>_<Description>` naming convention
- Use of `require` instead of `assert` (this repo uses `assert`)
- Tests in `package env` instead of `package env_test` (black-box)
- Missing `t.Helper()` in test helpers; missing `t.Cleanup()` for resource
  teardown

**Security**

- Secrets or credentials hard-coded
- Unvalidated external input

---

## YAML Review

Focus on `*.yaml` files.

**GitHub Actions Workflows**

- Triggered on broader events than necessary
- Action versions not pinned to a specific tag (avoid `@main` or `@latest`)
- Secrets referenced correctly and not exposed in logs
- Step outputs and env vars used correctly

**General YAML**

- Every document starts with `---` (required by `yamlfmt`
  `include_document_start: true`)
- No duplicate keys
- Anchors and aliases used correctly if present
