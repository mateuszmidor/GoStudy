default: proto

proto:
	cd sailworks/grpc/ && buf lint && buf generate --template=buf.gen.yaml && buf generate --template=buf.gen.tag.yaml

clean:
	rm -rf sailworks/grpc/gen

vet:
	go vet ./...

run: proto
	# run sailworks svc in background
	go build -o sailworks.bin sailworks/cmd/sailworks/*
	./sailworks.bin &

	# run shipyard application, it talks to sailworks and terminates
	go run shipyard/cmd/shipyard/*

	# stop sailworks svc
	pkill -f sailworks
	rm sailworks.bin

c4:
	cd docs/ && plantuml *.puml