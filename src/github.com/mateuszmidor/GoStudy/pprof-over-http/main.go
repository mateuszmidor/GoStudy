package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strings"
	"time"
)

func main() {
	runtime.SetBlockProfileRate(1)     // by default disabled
	runtime.SetMutexProfileFraction(1) // by default disabled
	go http.ListenAndServe("localhost:6060", http.DefaultServeMux)

	// keep the system under load, so there are always some cpu profile samples available
	start := time.Now()
	for time.Since(start).Minutes() < 10 {
		createLoad()
		time.Sleep(time.Second)
	}
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
