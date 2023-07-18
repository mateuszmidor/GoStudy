# bench-compare

Compare Go benchmark results from 2 different benchmark runs

# Run

```sh
make
```

```
go test -tags=slow -bench=. -benchtime=10000x -count=6 -benchmem > slow.txt
go test -tags=fast -bench=. -benchtime=10000x -count=6 -benchmem > fast.txt
go install golang.org/x/perf/cmd/benchstat@latest

benchstat slow.txt fast.txt
goos: linux
goarch: amd64
pkg: github.com/mateuszmidor/GoStudy/bench-compare
cpu: AMD Ryzen 7 4800H with Radeon Graphics
          │   slow.txt    │              fast.txt               │
          │    sec/op     │    sec/op     vs base               │
Concat-16   2145.50n ± 7%   13.54n ± 46%  -99.37% (p=0.002 n=6)

          │   slow.txt    │             fast.txt              │
          │     B/op      │    B/op     vs base               │
Concat-16   5344.000 ± 0%   4.000 ± 0%  -99.93% (p=0.002 n=6)

          │  slow.txt  │              fast.txt              │
          │ allocs/op  │ allocs/op   vs base                │
Concat-16   2.000 ± 0%   0.000 ± 0%  -100.00% (p=0.002 n=6)
```