---
description: Generate a PR title and body for the current branch (optional: pass base branch as $1, default main)
agent: plan
---

Generate a pull request title and body for the current branch against base
`${1:-main}`.

Rules:

- Title follows Conventional Commits (`type(scope): subject`, ≤72 chars)
- Body must have exactly these sections:

  ## Summary

  1-3 bullet points on intent and user-visible impact

  ## Changes

  Key changes, grouped logically (not a file-by-file recap)

  ## Testing

  How this was tested or what tests were added/updated

- Focus on _why_, not _what_
- Output ONLY the title and markdown body. Do not run gh pr create.

Commits in this branch: !`git log ${1:-main}..HEAD --oneline`

Full diff vs base: !`git diff ${1:-main}...HEAD`
