---
name: golang-validator
description: Run `golint ./...` and then `go vet ./...` in the current Go demo module to report Go style issues. Use when the user asks to validate Go code or find code problems in the current folder.
---

# Golint Runner

## Instructions
1. Confirm the current working directory is the Go demo module root (it should contain a `go.mod` file).
2. Run `pwd` from the current directory
3. Run `golint ./...` from the current directory.
4. Run `go vet ./...` from the current directory.
5. Return the raw output (and optionally a short summary if the output is long).

## Example
User request: "validate the current demo."

Agent behavior:
- Verify `go.mod` exists in the current folder
- Run `pwd`
- Show the results
- Run `golint ./...`
- Show the results
- Run `go vet ./...`
- Show the results

