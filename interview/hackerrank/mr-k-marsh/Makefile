
run: build
	./mr-k-marsh< TestCase11.txt

build: main.go
	go build -o mr-k-marsh main.go

profile: cpu.pprof
	go tool pprof -text -cum mr-k-marsh  cpu.pprof

profile-http: cpu.pprof
	go tool pprof -http=:8080 cpu.pprof