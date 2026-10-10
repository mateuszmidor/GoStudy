GOPATH ?= $(shell go env GOPATH)

.PHONY: run tools lint graph

run:
	go run ./cmd/calc

## Install the architecture linter (github.com/fe3dback/go-arch-lint)
tools:
	go install github.com/fe3dback/go-arch-lint@latest

## Check the code against the rules in .go-arch-lint.yml (needs `make tools` first)
lint:
	$(GOPATH)/bin/go-arch-lint check

## Regenerate go-arch-lint-graph.svg in the project root
graph:
	$(GOPATH)/bin/go-arch-lint graph
