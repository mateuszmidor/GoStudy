package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strings"

	"github.com/mateuszmidor/GoStudy/flight-finder/cmd/util"
)

func main() {
	// collect CPU profile
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	// collect traces
	traces, _ := os.Create("trace.out")
	defer traces.Close()
	trace.Start(traces)
	defer trace.Stop()

	runCLI()

	// collect memory profile
	heap, _ := os.Create("mem.out")
	defer heap.Close()
	runtime.GC() // get up-to-date statistics
	pprof.WriteHeapProfile(heap)
}

func runCLI() {
	const promptMsg = "Try: krk gdn. For exit: exit"
	const maxSegmentCount = 2
	fmt.Println(promptMsg)

	finder := util.NewConnectionFinder("../../segments.csv.gz", "../../airports.csv.gz", "../../nations.csv.gz", "\n")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "exit" {
			fmt.Println("exiting now.")
			break
		}

		if from, to, ok := parseFromTo(line); ok {
			fmt.Println("working...")
			finder.FindConnectionsAsText(os.Stdout, strings.ToUpper(from), strings.ToUpper(to), maxSegmentCount)
			fmt.Println()
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
