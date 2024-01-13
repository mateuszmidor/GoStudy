# WebServiceGin

Based on https://go.dev/doc/tutorial/web-service-gin

## Run

```sh
# run
go run .
```

## Test with curl

```sh
# list albums
curl http://localhost:8080/albums

# add new album
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

# get album by ID
curl http://localhost:8080/albums/2
```

## Test with [request.http](./request.http)

1. install VS Code plugin "REST Client"
1. click "Send Request"