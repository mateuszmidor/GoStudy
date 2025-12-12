default: proto

proto:
	cd sailworks/grpc/ && buf lint && buf generate --template=buf.gen.yaml && buf generate --template=buf.gen.tag.yaml
	cd sawmill/grpc/ && buf lint && buf generate --template=buf.gen.yaml && buf generate --template=buf.gen.tag.yaml

clean:
	rm -rf sailworks/grpc/gen
	rm -rf sawmill/grpc/gen

vet:
	go vet ./...

run: proto
	# run sailworks svc in background
	go build -o sailworks.bin sailworks/cmd/sailworks/*
	./sailworks.bin &

	# run sawmill svc in background
	go build -o sawmill.bin sawmill/cmd/sawmill/*
	./sawmill.bin &

	# run shipyard application, it talks to sailworks and sawmill and terminates
	go run shipyard/cmd/shipyard/*

	# stop services
	pkill -f sailworks
	pkill -f sawmill
	rm sailworks.bin sawmill.bin

c4:
	cd docs/ && plantuml *.puml