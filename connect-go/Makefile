default: proto
	go run .

proto: 
	buf lint || true # ignore lint warnings
	buf generate

clean:
	rm pingpong/*.go
	rm -rf pingpong/pingpongconnect