default: proto

proto:
	cd sawmill/grpc/ && buf lint && buf generate --template=buf.gen.yaml && buf generate --template=buf.gen.tag.yaml

clean:
	rm -rf sawmill/grpc/gen

vet:
	go vet ./...

run: proto
	# run sawmill svc in background
	go build -o sawmill.bin sawmill/cmd/sawmill/*
	./sawmill.bin &
	sleep 1

	# run shipyard application, it talks to sawmill and terminates
	go run shipyard/cmd/shipyard/*

	# stop services
	pkill -f sawmill
	rm sawmill.bin

c4:
	cd docs/ && plantuml *.puml