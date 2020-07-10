package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

func main() {
	// collect CPU profile
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	finder := newCliPathFinder("../../segments.csv.gz")
	runCLI(finder)

	// collect memory profile
	heap, _ := os.Create("mem.out")
	defer heap.Close()
	runtime.GC() // get up-to-date statistics
	pprof.WriteHeapProfile(heap)
}

func runCLI(f *cliPathFinder) {
	const promptMsg = "Try: krk gdn. For exit: exit"
	fmt.Println(promptMsg)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "exit" {
			return
		}

		if from, to, ok := parseFromTo(line); ok {
			fmt.Println("working...")
			start := time.Now()
			f.findConnections(strings.ToUpper(from), strings.ToUpper(to))
			d := time.Now().Sub(start)
			fmt.Printf("Took %dms\n", d.Milliseconds())
		} else {
			fmt.Println(promptMsg)
		}
	}
}

func parseFromTo(line string) (from string, to string, ok bool) {
	_, err := fmt.Sscanf(line, "%s %s\n", &from, &to)
	if err != nil {
		// fmt.Println(err)
		return "", "", false
	}
	return from, to, true
}
