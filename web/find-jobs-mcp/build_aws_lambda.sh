#!/usr/bin/env bash

set -euo pipefail

echo "Building AWS Lambda deployment package for GO..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap ./lambda

zip bootstrap.zip bootstrap
rm bootstrap

echo "Done. Upload bootstrap.zip in the AWS console."
