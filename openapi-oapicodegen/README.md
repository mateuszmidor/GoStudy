# OpenAPI with oapi-codegen

This demo shows how to generate Go code from OpenAPI 3.0.0 specs using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen).

## Install oapi-codegen

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

## Generate Go Code

**Generate client:**
```bash
./gen_client.sh
# Or manually:
oapi-codegen -package client -generate types,client fridge_api.yaml > generated_client/client.go
```

**Generate server:**
```bash
./gen_server.sh
# Or manually:
oapi-codegen -package server -generate types,std-http fridge_api.yaml > generated_server/server.go
```

## Run Server

```bash
go run ./cmd/server/main.go
```

The server will listen on port 8080.

## Test with curl

```bash
# Add a product
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Milk", "quantity": 2}'

# Get all products
curl http://localhost:8080/products

# Get a specific product
curl http://localhost:8080/products/Milk

# Withdraw 0.5 from product
curl -X PUT http://localhost:8080/products/Milk \
  -H "Content-Type: application/json" \
  -d '{"quantity": 0.5}'
```

## Generated Code

- `generated_client/client.go` - Go client library with types and HTTP client
- `generated_server/server.go` - net/http server with server interface and types
- `cmd/server/main.go` - Example server implementation using net/http ServeMux

## Key Differences from openapi-generator

- Single Go binary (no Docker required)
- Uses Go standard library net/http (no external router dependency)
- Types, client, and server in single files
- More lightweight output