TAG=mateuszmidor/delve:1.21

default:
	docker build . -t ${TAG}

push:
	docker login
	docker push ${TAG}