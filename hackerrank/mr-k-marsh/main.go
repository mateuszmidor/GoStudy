package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"sort"
)

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

func canFenceRect(limits [][]RunLength, x, y, w, h int) bool {
	if limits[y][x].cleanX < w ||
		limits[y][x].cleanY < h ||
		limits[y+h-1][x].cleanX < w ||
		limits[y][x+w-1].cleanY < h {
		return false
	}
	return true
}

type Field struct {
	sizex, sizey, perimeter int
}

type RunLength struct {
	cleanX int // how many fence posts are possible going to the right
	cleanY int // how many fence posts are possible going to the bottom
}

func newRunLength(w, h int) [][]RunLength {
	m := make([][]RunLength, h)
	for y := 0; y < h; y++ {
		m[y] = make([]RunLength, w)
	}
	return m
}

func computeRunLength(runLength [][]RunLength, grid []string, w, h int) {
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
}

func getMaxPerimeter(grid []string, w, h int) int { // -1 means no rectangle found
	if w < 2 || h < 2 || w > 500 || h > 500 {
		return -1
	}

	sizes := []Field{}
	for sizey := 2; sizey <= h; sizey++ {
		for sizex := 2; sizex <= w; sizex++ {
			sizes = append(sizes, Field{sizex, sizey, calcPerimeter(sizex, sizey)})
		}
	}

	compare := func(i, j int) bool {
		return (sizes[i].perimeter >= sizes[j].perimeter)
	}
	sort.Slice(sizes, compare)

	runLength := newRunLength(w, h)
	computeRunLength(runLength, grid, w, h)
	for _, size := range sizes {
		numStepsY := h - size.sizey + 1
		numStepsX := w - size.sizex + 1
		for y := 0; y < numStepsY; y++ {
			for x := 0; x < numStepsX; x++ {
				if canFenceRect(runLength, x, y, size.sizex, size.sizey) {
					return size.perimeter
				}

			}
		}
	}
	return -1
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
