#!/bin/bash

# dont run tests, run single benchmark
go test -run=NONE -bench=BenchmarkAlloc -memprofile=mem.out

# print memprofile in form of graphviz svg
go tool pprof -web ./testexamples.test mem.out