#!/usr/bin/env bash

function die() {
    echo "Error: $1. Exiting"
    exit 1
}

echo "Building AWS Lambda deployment package for GO..."
go build hello.go || die "Build failed"
zip hello.zip hello 
echo "Done"