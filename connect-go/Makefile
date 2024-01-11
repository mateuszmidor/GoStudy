default: proto
	go run .

proto: 
	buf lint || true # ignore lint warnings
	buf generate --template=buf.gen.yaml
	buf generate --template=buf.gen.tag.yaml

clean:
	rm -rf gen