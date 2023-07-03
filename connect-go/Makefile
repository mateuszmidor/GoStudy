default: proto-to-connect-go
	go run .

proto-to-connect-go: 
	buf lint || true # ignore lint warnings
	buf generate

clean:
	rm pingpong/*.go
	rm -rf pingpong/pingpongconnect