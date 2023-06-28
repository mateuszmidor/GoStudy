default: proto-to-go
	go run .

proto-to-go: 
	protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative pingpong/*.proto

clean:
	rm pingpong/*.go