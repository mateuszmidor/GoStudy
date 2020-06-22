package main

import (
	"image"
	"math"
	"sync"
)

const xmin, ymin, xmax, ymax = -2, -2, +2, +2

func goMandel(img *image.RGBA, numParallel int) {
	processParallelInSegments(numParallel, img)
}

func processParallelInSegments(numSegments int, img *image.RGBA) {
	var wg sync.WaitGroup
	var ribbonHeight = height / numSegments

	for i := 0; i < numSegments; i++ {
		wg.Add(1)
		go func(nSegment int) {
			beginy := nSegment * ribbonHeight
			endy := (nSegment + 1) * ribbonHeight
			processSegment(beginy, endy, img)
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

			c := mandelbrot(x, y)
			i := img.PixOffset(px, py)
			s := img.Pix[i : i+4 : i+4]
			s[0] = c
			s[1] = c
			s[2] = c
			s[3] = 255
		}
	}
}

func mandelbrot(x, y float64) uint8 {
	const iterations = 200
	const contrast = 15

	var imag float64
	var real float64

	for n := uint8(0); n < iterations; n++ {
		i := imag
		r := real

		real = (r*r - i*i) + x
		imag = (r*i + i*r) + y

		if math.Sqrt(imag*imag+real*real) > 2.0 {
			return 255 - contrast*n
		}
	}
	return 0
}
