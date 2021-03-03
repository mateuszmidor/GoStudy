# Swagger for GO using goswagger

<https://github.com/go-swagger/go-swagger>

## Swagger 2.0/OpenAPI 3.0 YAML online editors with live preview

- https://app.swaggerhub.com/ - needs free account
- https://editor.swagger.io/ - needs no account

## Install goswagger

```bash
dir=$(mktemp -d) 
git clone https://github.com/go-swagger/go-swagger "$dir" 
cd "$dir"
go install ./cmd/swagger
```

## Validate spec

```bash
swagger validate spec.json
> The swagger spec at "spec.json" is valid against swagger specification 2.0
```

## Serve API as http

```bash
swagger serve spec.json
# web browser opens up
```

## Generate client (in folder with go.mod)

```bash
swagger generate client -f spec.json
 # lots of output, new folders created: "client" and "models"
 go mod tidy # download dependencies
 ```

## Generate server (in folder with go.mod)

```bash
swagger generate server -f spec.json
 # lots of output, new folders created: "cmd", "models" and "restapi"
 go mod tidy # download dependencies
```

## Generate just the data model (in folder with go.mod)

```bash
swagger generate model --spec spec.json
 # lots of output, new folder created: "models"
```

## Generate markdown documentation for spec

```bash
swagger generate markdown -f spec.json --output spec.md
 # new file created: "spec.md"
```
