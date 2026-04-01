#!/bin/bash
set -e

oapi-codegen -package client -generate types,client fridge_api.yaml > generated_client/client.go
echo "Client generated in generated_client/client.go"