default: proto

proto:
	cd shipyard/api/ && buf lint && buf generate --template=buf.gen.yaml && buf generate --template=buf.gen.tag.yaml

clean:
	rm -rf shipyard/api/gen/

run: proto
	# run sailworks svc in background
	go build -o sailworks shipyard/cmd/sailworks/*
	./sailworks &

	# run shipyard application, it talks to sailworks and terminates
	go run shipyard/cmd/shipyard/*

	# stop sailworks svc
	pkill -f sailworks
	rm sailworks
