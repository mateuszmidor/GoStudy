package main

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 900, 600
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

func plotSVG3D(w io.Writer, formula string) error {
	// create eval
	eval, err := NewEvalXY(formula)
	if err != nil {
		return err
	}

	// test eval
	_, err = eval.Eval(0, 0)
	if err != nil {
		return fmt.Errorf("Evaluation error in %q: %s", formula, err)
	}

	// do actual plot
	doPlot(w, eval)

	return nil
}

func doPlot(w io.Writer, eval MathEvalXY) {
	writeSVHHeader(w)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, eval)
			bx, by := corner(i, j, eval)
			cx, cy := corner(i, j+1, eval)
			dx, dy := corner(i+1, j+1, eval)
			writeSVGPolygon(w, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	writeSVGFooter(w)
}

func corner(i, j int, eval MathEvalXY) (float64, float64) {
	// for isometric cast
	var sin30, cos30 = math.Sin(angle), math.Cos(angle)

	// find point (x,y) in corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// calc the height at (x,y)
	z, err := eval.Eval(x, y)
	if err != nil {
		z = 0.0
	}

	// isometric cast to 2D canvas
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}
