// сервер позволяет рисовать фигуры лиссажоу с параметрами, задавая их так:
// http://192.168.1.33:8000/lis/?cycles=10
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
	"net/url"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // Первый цвет палитры
	blackIndex = 1 // Следующий цвет палитры
)

func lissajous(out io.Writer, cyclesCount int) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)
	if cyclesCount <= 0 || cyclesCount > 30 {
		cyclesCount = cycles
	}
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	freq := rand.Float64() * 3.0 // относительная частота колебаний Y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cyclesCount)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
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

func handler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	queryParams := u.Query()
	fmt.Println("Params:", queryParams)
	cycles := queryParams.Get("cycles")
	fmt.Println("cycles:", cycles)
	cyclesCount, err := strconv.Atoi(cycles)
	if err != nil {
		fmt.Println("conversion failsed:", err)
		return
	}

	lissajous(w, cyclesCount)
}

func main() {
	http.HandleFunc("/lis/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
