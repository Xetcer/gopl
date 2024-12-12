package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

type Func func(complex128) complex128

var palettes = []color.RGBA{
	{170, 57, 57, 255},
	{170, 108, 57, 255},
	{34, 102, 102, 255},
	{45, 136, 45, 255},
}

var chosenColors = map[complex128]color.RGBA{}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px, py, z4(z))
		}
	}
	f, err := os.Create("maldebrot.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := png.Encode(f, img); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}
	// err := png.Encode(os.Stdout, img)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func z4(z complex128) color.Color {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	fPrime := func(z complex128) complex128 {
		return (z - 1/(z*z*z)) / 4
	}
	return newton(z, f, fPrime)
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}

func newton(z complex128, f, fPrime Func) color.Color {
	const iterations = 37
	for i := uint8(0); i < iterations; i++ {
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := chosenColors[root]
			if !ok {
				if len(palettes) == 0 {
					panic("no colors left")
				}
				c = palettes[0]
				palettes = palettes[1:]
				chosenColors[root] = c
			}
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i) / math.Log(iterations))
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			colorRGBA := 255 - contrast*n
			return color.RGBA{255, 255 - n, colorRGBA, 255}
			// return color.RGBA{colorRGBA, colorRGBA, colorRGBA, 255}
			// return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
