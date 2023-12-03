# run in Docker as non-root user
default: build run
	

build: 
	docker build . -t embedded-postgress-demo

run:
	docker run --rm -it --user regular-user embedded-postgress-demo 