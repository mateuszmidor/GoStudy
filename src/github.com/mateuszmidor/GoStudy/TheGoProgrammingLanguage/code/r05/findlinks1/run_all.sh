#!/bin/bash

# fetch html and pass it to program stdin
go get golang.org/x/net/html && wget -O - --quiet golang.org | go run .