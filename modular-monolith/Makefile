default: proto

proto:
	cd api/ && buf lint && buf generate --template=buf.gen.yaml && buf generate --template=buf.gen.tag.yaml

clean:
	rm -rf api/gen/

run: proto
	# run sailworks svc in background
	go build -o sailworks cmd/sailworks/*
	./sailworks &

	# run shipyard application, it talks to sailworks and terminates
	go run cmd/shipyard/*

	# stop sailworks svc
	pkill -f sailworks
	rm sailworks
