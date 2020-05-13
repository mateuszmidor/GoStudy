// Project: go-routined mandelbrot, up to 3x speedup on 4 core
// Usage: go run . ;  firefox mandel.png
package main

// #cgo CPPFLAGS: -DCGO=1
// #include <stdint.h>
// void freeMandel(char* data);
// char* makeMandel(size_t width, size_t height, size_t numParallel);
// void makeMandel2(size_t width, size_t height, char* data, size_t numParallel) ;
import "C"

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	size                   = 1024
	width, height          = size, size
)

func main() {
	println(runtime.GOMAXPROCS(1))
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	benchCGOMandel(img)
	saveImage(img, "cgo.png")

	// benchGOMandel(img)
	// saveImage(img, "go.png")
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

// Variant 1: C side allocates memory
func cgoMandel(img *image.RGBA, numParallel int) {
	// get C memory allocated from C function
	crgba := C.makeMandel(width, height, C.size_t(numParallel))

	// convert to slice of bytes
	rgba := C.GoBytes(unsafe.Pointer(crgba), width*height*4)

	// copy bytes to image
	copy(img.Pix, rgba)

	// release C memory
	C.freeMandel(crgba)
}

// Variant 2: GO side allocates memory
func cgoMandel2(img *image.RGBA, numParallel int) {
	// cast GO slice of bytes to char*
	crgba := (*C.char)(unsafe.Pointer(&img.Pix[0]))

	// run C function to fill the image
	C.makeMandel2(width, height, crgba, C.size_t(numParallel))
}

/*
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
*/

func printTimes(times []time.Duration) {
	for n, time := range times {
		printTime(n, time)
	}
}

func printTime(numWorkers int, time time.Duration) {
	fmt.Printf("%4.4s - %0.2vms\n", strconv.Itoa(numWorkers), time)
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
