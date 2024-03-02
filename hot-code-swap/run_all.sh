#!/usr/bin/env bash

echo "Installing 'reflex', it gets in $GOPATH/bin"
go get github.com/cespare/reflex

echo "Running 'reflex', make changes to any .go file and see the program gets re-run"
$GOPATH/bin/reflex -r '\.go$' go run .