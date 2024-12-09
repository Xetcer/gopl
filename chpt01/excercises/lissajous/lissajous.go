/*
go run lissajous.go web  - http://localhost:8000
вывод анимации в браузер фигуры лисажоу
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{R: 0, G: 0xff, B: 0, A: 0xff}, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff}, color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff}}

const (
	blackIndex = 0 // Первый цвет палитры
	greenIndex = 1 // Следующий цвет палитры
	redIndex   = 2 // Следующий цвет палитры
	blueIndex  = 3 // Следующий цвет палитры
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	// rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe(":8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	freq := rand.Float64() * 3.0 // относительная частота колебаний Y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	colorIndex := uint8(greenIndex)
	colorIndex = uint8(rand.Intn(3) + 1)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Println(err)
		return
	}
}
