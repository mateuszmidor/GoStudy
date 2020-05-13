// Project: go-routined mandelbrot, up to 3x speedup on 4 core
// Usage: go run . > mandel.png;  firefox mandel.png
package main

import (
	"fmt"
	"image"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	size                   = 1 * 1024
	width, height          = size, size
)

func main() {
	fmt.Printf("g times for parallel mandel %dx%d:\n", width, height)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	numSegments := 1
	for numSegments < height {
		start := time.Now()
		processParallelInSegments(numSegments, img)
		duration := time.Now().Sub(start)

		printTime(numSegments, duration)
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

			c := mandelbrot(z)
			i := img.PixOffset(px, py)
			s := img.Pix[i : i+4 : i+4]
			s[0] = c
			s[1] = c
			s[2] = c
			s[3] = 255
		}
	}
}

func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2.0 {
			return 255 - contrast*n
		}
	}
	return 0
}

func printTime(numWorkers int, time time.Duration) {
	fmt.Fprintf(os.Stderr, "%4.4s - %0.2dms\n", strconv.Itoa(numWorkers), time.Milliseconds())
}
