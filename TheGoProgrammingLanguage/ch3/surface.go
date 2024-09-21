package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = 0.4 * height
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("An error occured, %v", err)
		panic(err)
	}
}

func main() {
	f, err := os.Create("surface.html")
	checkErr(err)

	defer func() {
		err = f.Close()
		checkErr(err)
	}()

	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, "<!DOCTYPE html>\n<html>\n<body>\n")
	checkErr(err)
	_, err = fmt.Fprintf(w, "<svg xmlns='https://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>\n", width, height)

	checkErr(err)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			_, err = fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			checkErr(err)
		}
	}
	_, err = fmt.Fprintln(w, "</svg>")
	checkErr(err)
	_, err = fmt.Fprintf(w, "</body>\n</html>\n")
	checkErr(err)
	w.Flush()
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
