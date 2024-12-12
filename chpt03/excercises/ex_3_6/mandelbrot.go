package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		espX                   = (xmax - xmin) / width
		espY                   = (ymax - ymin) / height
	)

	offX := []float64{-espX, espX}
	offY := []float64{-espY, espY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			subpixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					subpixels = append(subpixels, mandelbrot(z))
				}
			}
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px, py, supersampling(subpixels))
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

func supersampling(subpixels []color.Color) color.Color {
	r, g, b := 0, 0, 0
	for _, rgba := range subpixels {
		R, G, B, _ := rgba.RGBA()
		r += int(R)
		g += int(G)
		b += int(B)
	}
	return color.RGBA{R: uint8(r / 4), G: uint8(g / 4), B: uint8(b / 4), A: 255}
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

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
