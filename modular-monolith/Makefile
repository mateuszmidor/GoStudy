default: proto

proto:
	cd shipyard/modules/sailworks/grpc/ && buf lint && buf generate --template=buf.gen.yaml && buf generate --template=buf.gen.tag.yaml

clean:
	rm -rf shipyard/modules/sailworks/grpc/gen

run: proto
	# run sailworks svc in background
	go build -o sailworks shipyard/cmd/sailworks/*
	./sailworks &

	# run shipyard application, it talks to sailworks and terminates
	go run shipyard/cmd/shipyard/*

	# stop sailworks svc
	pkill -f sailworks
	rm sailworks
