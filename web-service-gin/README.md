# WebServiceGin

Based on https://go.dev/doc/tutorial/web-service-gin

## Run

```sh
# run
go run .

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