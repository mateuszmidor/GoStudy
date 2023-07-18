N = 10000x
NUM_SAMPLES = 6

default:
	go test -tags=slow -bench=. -benchtime=$(N) -count=$(NUM_SAMPLES) -benchmem > slow.txt
	go test -tags=fast -bench=. -benchtime=$(N) -count=$(NUM_SAMPLES) -benchmem > fast.txt

	go install golang.org/x/perf/cmd/benchstat@latest
	@echo
	benchstat slow.txt fast.txt