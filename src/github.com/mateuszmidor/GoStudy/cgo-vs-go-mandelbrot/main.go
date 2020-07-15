// Project: go-routined mandelbrot, up to 3x speedup on 4 core
// Usage: go run . ;  firefox mandel.png
package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	size          = 1024 * 4
	width, height = size, size
)

func main() {
	// tracef, _ := os.Create("trace.out")
	// defer tracef.Close()
	// trace.Start(tracef)
	// defer trace.Stop()

	cgoMandelImg := image.NewRGBA(image.Rect(0, 0, width, height))
	benchCGOMandel(cgoMandelImg)
	saveImage(cgoMandelImg, "cgo-mandel.png")

	println()

	goMandelImg := image.NewRGBA(image.Rect(0, 0, width, height))
	benchGOMandel(goMandelImg)
	saveImage(goMandelImg, "go-mandel.png")
}

func benchCGOMandel(img *image.RGBA) {
	fmt.Printf("cgo times for parallel mandel %dx%d:\n", width, height)
	for numSegments := 1; numSegments < height; numSegments *= 2 {
		start := time.Now()
		cgoMandel2(img, numSegments)
		duration := time.Now().Sub(start)
		printTime(numSegments, duration)
	}
}

func benchGOMandel(img *image.RGBA) {
	fmt.Printf("go times for parallel mandel %dx%d:\n", width, height)
	for numParallel := 1; numParallel <= runtime.NumCPU(); numParallel++ {
		start := time.Now()
		goMandel(img, numParallel)
		duration := time.Now().Sub(start)
		printTime(numParallel, duration)
	}
}

func printTime(numWorkers int, time time.Duration) {
	fmt.Printf("%4.4s - %0.2dms\n", strconv.Itoa(numWorkers), time.Milliseconds())
}

func saveImage(img image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("could not create file: %v\n", err)
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		fmt.Printf("could not encode png: %v\n", err)
		return
	}
}
