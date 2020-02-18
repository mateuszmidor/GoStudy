// Project: go-routined mandelbrot, up to 3x speedup on 4 core
// Usage: go run . > mandel.png && firefox mandel.png
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	numSegments := 1
	for numSegments < height {
		start := time.Now()
		processParallelInSegments(numSegments, img)
		duration := time.Now().Sub(start)
		fmt.Fprintf(os.Stderr, "Mandelbrot %dx%d in %d goroutines generated in %v\n", width, height, numSegments, duration)
		numSegments *= 2
	}

	png.Encode(os.Stdout, img) // ignoring erors
}

func processParallelInSegments(numSegments int, img *image.RGBA) {
	var wg sync.WaitGroup
	for i := 0; i < numSegments; i++ {
		wg.Add(1)
		go func(nSegment int) {
			processSegment(nSegment*height/numSegments, (nSegment+1)*height/numSegments, img)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func processSegment(beginy, endy int, img *image.RGBA) {
	for py := beginy; py < endy; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2.0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
