package main

import (
	"os"
	"runtime/trace"
	"strings"
)

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()
	createLoad()
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
