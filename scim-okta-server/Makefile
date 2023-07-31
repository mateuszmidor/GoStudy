default:
	@echo "specify target: linux or macos"
	
linux:
	echo "not implemented yet"
	exit 1
	# docker build -t lt .
	# docker run -it --rm --name=lt --net=host lt --subdomain dob-mateusz-scim --port 33000

# for some reason, localtunnel run from docker on mac, times out with http 504
macos:
	go run *.go &
	brew install localtunnel
	lt --subdomain dob-mateusz-scim --port 33000 