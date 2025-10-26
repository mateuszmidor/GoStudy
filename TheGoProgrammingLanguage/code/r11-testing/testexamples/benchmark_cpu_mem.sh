#!/bin/bash

# run only benchmarks, no regular tests
go test -v -run=NONE -bench="BenchmarkAfter|BenchmarkAlloc" -benchmem  