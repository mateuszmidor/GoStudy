# OpenAPI with openapi-generator

This demo shows how to generate Go code from OpenAPI 3.0.0 specs using [openapi-generator](https://openapi-generator.tech/).

## Install openapi-generator (Docker)

```bash
# Docker image is used for code generation
docker pull openapitools/openapi-generator-cli:latest
```
<!--
## Generate Go Code -->

**Generate client:**
```bash
./gen_client.sh
# Or manually:
docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate \
  -i /local/fridge_api.yaml \
  -g go \
  -o /local/generated_client \
  --additional-properties=packageName=client,moduleName=github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_client
```

**Generate server:**
```bash
./gen_server.sh
# Or manually:
docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate \
  -i /local/fridge_api.yaml \
  -g go-server \
  -o /local/generated_server \
  --additional-properties=packageName=server,useGinFramework=false,moduleName=github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_server
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

- `generated_client/` - Go client library with types and HTTP client
- `generated_server/go/` - Go server with types, router, and service interface
- `cmd/server/main.go` - Example server implementation using generated server code

## Key Differences from oapi-codegen
