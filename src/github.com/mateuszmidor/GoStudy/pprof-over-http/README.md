# pprof example: live cpu and memory profile over http

(No convenient run_all.sh since this is fully interactive example)

## Run

```bash
go run . & # this runs http server on port 6060; automatically terminates after 10 min
```

## See all profiles

```bash
firefox http://localhost:6060/debug/pprof/
```

## CPU profile

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=30 # beautiful web ui
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30             # console ui
```

## Memory profile

```bash

go tool pprof http://localhost:6060/debug/pprof/heap   # inuse_space, can use -http=:8080
go tool pprof http://localhost:6060/debug/pprof/allocs # alloc_space, can use -http=:8080
```

## Contended mutex holders

```bash
go tool pprof http://localhost:6060/debug/pprof/mutex # can use -http=:8080
```

## Goroutine blocking profile

```bash
go tool pprof http://localhost:6060/debug/pprof/block # can use -http=:8080
```

## Example result of memory profile: alloc_space

![Call graph with allocations](heap.gif)
