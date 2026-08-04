.PHONY: client server run

client:
	oapi-codegen -package client -generate types,client fridge_api.yaml > generated_client/client.go
	@echo "Client generated in generated_client/client.go"

server:
	oapi-codegen -package server -generate types,std-http fridge_api.yaml > generated_server/server.go
	@echo "Server generated in generated_server/server.go"

run:
	go run .
