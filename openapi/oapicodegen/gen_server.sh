#!/bin/bash
set -e

oapi-codegen -package server -generate types,std-http fridge_api.yaml > generated_server/server.go
echo "Server generated in generated_server/server.go"