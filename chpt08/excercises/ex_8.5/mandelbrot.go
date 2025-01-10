package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	var wg sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// создаем канал для передачи номеров строк между горутинами
	size := make(chan int, height)
	go func() {
		for row := 0; row < height; row++ {
			size <- row
		}
		close(size) // закроем канал когда отправим все строки
	}()

	for i := 0; i < 2; i++ { // создаем 2 горутины
		wg.Add(1)
		go func() {
			for py := range size { // получаем номера строк из канала
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Точка (px, py) представляет комплексное значение z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

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

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
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
