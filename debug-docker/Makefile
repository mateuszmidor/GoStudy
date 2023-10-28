IMAGE_NAME=debug-docker

default: image run

image:
	docker build . -t ${IMAGE_NAME}

run:
	docker run -t --rm --name debug-docker ${IMAGE_NAME}

debug:
	docker run -t --rm --name debug-docker -p 2345:2345 ${IMAGE_NAME} /go/bin/dlv --continue --listen=:2345 --headless=true --accept-multiclient --api-version=2 exec /app/HelloWorld