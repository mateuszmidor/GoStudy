#!/usr/bin/env bash

echo "Installing gomock module"
go get -u github.com/golang/mock/gomock

echo "Installing mockgen interface mock generator"
go get -u github.com/golang/mock/mockgen
go install github.com/golang/mock/mockgen # go/bin must be on PATH to use it

echo "Generating mocks for interfaces"
go generate -v .

echo "Running tests"
go test . -v