#!/usr/bin/env bash

echo "Installing gomock module"
go get -u github.com/golang/mock/gomock

echo "Installing mockgen interface mock generator"
go get -u github.com/golang/mock/mockgen

cd src/

echo "Generating mocks for interfaces"
go generate -v .

echo "Running tests"
go test .