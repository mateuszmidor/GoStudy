linux:
	docker build -t lt .
	docker run -it --rm --name=lt --net=host lt --subdomain dob-mateusz-test --port 33000

# for some reason, localtunnel run from docker on mac, times out with http 504
macos:
	brew install localtunnel
	lt --subdomain dob-mateusz-test --port 33000