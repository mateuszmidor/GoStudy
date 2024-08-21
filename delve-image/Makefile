TAG=mateuszmidor/delve:1.22

default:
	docker build . -t ${TAG}

push:
	docker login
	docker push ${TAG}