#!/bin/bash

# dont run tests, run single benchmark
go test -run=NONE -bench=BenchmarkFibb -cpuprofile=cpu.out

# print cpuprofile in form of ACII table
go tool pprof -text ./testexamples.test cpu.out