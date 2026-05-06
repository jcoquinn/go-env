---
description: Generate a Conventional Commits message for staged changes
agent: plan
---

Write a commit message for the staged changes below.

Rules:

- Format: `type(scope): subject` — imperative mood, ≤72 chars, no trailing
  period
- Valid types: feat, fix, refactor, test, docs, chore, perf, ci, build, style
- Omit scope if it adds no information
- Blank line between subject and body
- Body wrapped at 72 chars; explain the _why_, not the _what_
- Use `type!` and a `BREAKING CHANGE:` footer for breaking changes
- Output ONLY the commit message in a fenced code block. Do not run git commit.

Staged diff: !`git diff --cached`

Recent commits for style reference: !`git log -10 --oneline`
