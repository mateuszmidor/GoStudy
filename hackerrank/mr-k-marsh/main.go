package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"

	"sort"
)

type Field struct {
	width, height, perimeter int
}

type RunLength struct {
	cleanX int // how many fence posts are possible going to the right
	cleanY int // how many fence posts are possible going to the bottom
}

/*
 * Complete the 'kMarsh' function below.
 *
 * The function accepts STRING_ARRAY grid as parameter.
 */

func calcPerimeter(w, h int) int {
	if w < 2 || h < 2 {
		return -1
	}

	return 2*(w-1) + 2*(h-1)
}

func canFenceRect(runLength [][]RunLength, x, y, w, h int) bool {
	if runLength[y][x].cleanX < w ||
		runLength[y][x].cleanY < h ||
		runLength[y+h-1][x].cleanX < w ||
		runLength[y][x+w-1].cleanY < h {
		return false
	}
	return true
}

func newRunLength(w, h int) [][]RunLength {
	m := make([][]RunLength, h)
	for y := 0; y < h; y++ {
		m[y] = make([]RunLength, w)
	}
	return m
}

func computeRunLength(grid []string, w, h int) [][]RunLength {
	runLength := newRunLength(w, h)
	for y := 0; y < h; y++ {
		var count int
		for x := w - 1; x >= 0; x-- {
			if grid[y][x] == 'x' {
				count = 0
			} else {
				count++
			}
			runLength[y][x].cleanX = count
		}
	}

	for x := 0; x < w; x++ {
		var count int
		for y := h - 1; y >= 0; y-- {
			if grid[y][x] == 'x' {
				count = 0
			} else {
				count++
			}
			runLength[y][x].cleanY = count
		}
	}

	return runLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sortFieldsByPerimeterDescending(fields []Field) {
	compare := func(i, j int) bool {
		return (fields[i].perimeter >= fields[j].perimeter)
	}
	sort.Slice(fields, compare)
}

func getMaxPerimeter(grid []string, w, h int) int { // -1 means no rectangle found
	if w < 2 || h < 2 || w > 500 || h > 500 {
		return -1
	}

	// prepare time for 500x500 input is below 200ms
	start := time.Now()
	fields := generateAllPossibleFieldSizes(h, w)
	sortFieldsByPerimeterDescending(fields)
	runLength := computeRunLength(grid, w, h)
	fmt.Println("Precompute time:", time.Since(start))

	for _, field := range fields {
		numStepsY := h - field.height + 1
		numStepsX := w - field.width + 1
		for y := 0; y < numStepsY; y++ {

			x := 0 // .x...
			for {
				if x >= numStepsX {
					break
				}
				// clear := runLength[y][x].cleanX
				// if x < numStepsX && clear < field.width {
				// 	x += clear + 1
				// 	continue
				// }
				// if x >= numStepsX {
				// 	break
				// }

				if canFenceRect(runLength, x, y, field.width, field.height) {
					return field.perimeter
				}

				x += 1
			}
		}
	}
	return -1
}

func generateAllPossibleFieldSizes(h int, w int) []Field {
	fields := []Field{}
	for sizey := 2; sizey <= h; sizey++ {
		for sizex := 2; sizex <= w; sizex++ {
			fields = append(fields, Field{sizex, sizey, calcPerimeter(sizex, sizey)})
		}
	}
	return fields
}

func kMarsh(grid []string) {
	max := getMaxPerimeter(grid, len(grid[0]), len(grid))
	// Print the result
	if max > -1 {
		fmt.Println(max)
	} else {
		fmt.Println("impossible")
	}
}

func main() {
	f, perr := os.Create("./cpu.pprof")
	if perr != nil {
		log.Fatal(perr)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	_, err = strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)

	var grid []string

	for i := 0; i < int(m); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	kMarsh(grid)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
