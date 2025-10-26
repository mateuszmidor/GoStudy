#!/bin/bash

# dont run tests, run single benchmark
go test -run=NONE -bench=BenchmarkAfter -blockprofile=block.out

# print blockprofile in form of graphviz svg
go tool pprof -web ./testexamples.test block.out