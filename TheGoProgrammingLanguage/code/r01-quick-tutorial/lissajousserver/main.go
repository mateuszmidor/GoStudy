// Project: lissajousserver; curves (the ones you can see on oscilloscope) served over http!
// Usage:
//	./lissajousserver
//	firefox localhost:8000/curves?size=400
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/curves", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var sizeInt int
	sizeStr := r.URL.Query().Get("size")
	sizeInt, err := strconv.Atoi(sizeStr)
	if err != nil {
		sizeInt = 200
	}

	lissajous(w, sizeInt)
}

func lissajous(out io.Writer, size int) {
	const (
		cycles  = 5
		res     = 0.001
		nframes = 64
		delay   = 8
	)
	var palette = []color.Color{color.White, color.Black}

	const whiteIndex = 0
	const blackIndex = 1

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // ignoring EncodeAll errors
}
