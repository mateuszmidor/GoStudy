#!/bin/bash

# collect coverage info
go test -v -coverprofile=coverage.out -covermode=count

# display coverage in html form
go tool cover -html=coverage.out