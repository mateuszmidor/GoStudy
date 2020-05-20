package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strings"
)

func main() {
	// trace
	traces, _ := os.Create("traces.out")
	defer traces.Close()
	trace.Start(traces)
	defer trace.Stop()

	// collect CPU profile
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	// do actual work
	createLoad()

	// collect memory profile
	heap, _ := os.Create("heap.out")
	defer heap.Close()
	runtime.GC() // get up-to-date statistics
	pprof.WriteHeapProfile(heap)
}

func createLoad() int {
	return calcPrefix() + calcPostfix()
}

func calcPrefix() int {
	const descr = "descrIption"
	c := 0
	for i := 0; i < 12*1000*1000; i++ {
		if strings.Contains(strings.ToLower(descr), "desc") {
			c++
		}
	}
	return c
}

func calcPostfix() int {
	const descr = "descrIption"
	c := 0
	for i := 0; i < 4*1000*1000; i++ {
		if strings.Contains(strings.ToLower(descr), "tion") {
			c++
		}
	}
	return c
}
