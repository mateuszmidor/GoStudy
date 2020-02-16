package main

import (
	"fmt"
	"io"
)

func writeSVHHeader(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
}

func writeSVGPolygon(w io.Writer, ax, ay, bx, by, cx, cy, dx, dy float64) {
	fmt.Fprintf(w, "<polygon style='stroke: red' points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
}

func writeSVGFooter(w io.Writer) {
	fmt.Fprintf(w, "</svg>")
}
