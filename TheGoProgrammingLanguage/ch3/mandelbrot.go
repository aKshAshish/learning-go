package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	f, err := os.Create("mandelbrot.png")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(f, img)
}

/*
func oversample(x, y float64) color.Color {
	z1 := mandelbrot(complex(x+1, y))
	z2 := mandelbrot(complex(x, y))
	z3 := mandelbrot(complex(x, y+1))
	z4 := mandelbrot(complex(x+1, y+1))
	val := (z1 + z2 + z3 + z4) / 4
	if val == 0 {
		return color.Black
	}
	return color.Gray{val}
}
*/

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 255 - contrast*n, 255 - contrast*n, 255 - contrast*n}
		}
	}
	return color.Black
}
