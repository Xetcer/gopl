/*
Exercise 3.3
Surface вычисляет SVG-представление трехмерного графика функции.
Окрасьте каждый многоугольник цветом, зависящим от его высоты, так,
чтобы пики были красными (#ff0000), а низины — синими (#0000ff).

http://localhost:8000/svg/?height=350
*/

package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// width, height = 600, 320            // Размер канвы в пикселях
	cells   = 100 // Количество ячеек сетки
	xyrange = 30  // Диапазон осей (-xyrange..+xyrange)
	// xyscale       = width / 2 / xyrange // Пикселей в единице x или y
	// zscale = height * 0.4 // Пикселей в единице z
	angle = math.Pi / 6 // Углы осей x, y (=30°)
)

var width, height = 600, 320
var xyscale = width / 2 / xyrange  // Пикселей в единице x или y
var zscale = 0.4 * float64(height) // Пикселей в единице z

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func handler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	queryParams := u.Query()
	height, err = strconv.Atoi(queryParams.Get("height"))
	if err != nil {
		height = 320
	}
	fmt.Println("qHeight set:", height)
	width, err = strconv.Atoi(queryParams.Get("width"))
	if err != nil {
		width = 600
	}
	fmt.Println("qWidth set:", width)

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg'> "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// exercise 3.1
			ax, ay, ct, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, ct1, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, ct2, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, ct3, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			// exercise 3.3c
			var color string

			switch {
			case ct == 1 || ct1 == 1 || ct2 == 1 || ct3 == 1:
				color = "#ff0000"
			case ct == 2 || ct1 == 2 || ct2 == 2 || ct3 == 2:
				color = "#0000ff"
			default:
				color = "#00ff00"
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"fill='#222222' stroke='%s' stroke-width='0.4'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func main() {
	http.HandleFunc("/svg/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func corner(i, j int) (float64, float64, int, error) {
	// Ищем угловую точку (x,y) ячейки (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z, ct := f(x, y)
	// exercise 3.1
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, fmt.Errorf("invalid value")
	}

	// Изометрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := float64(width/2) + (x-y)*cos30*float64(xyscale)
	sy := float64(height/2) + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy, ct, nil

}

func f(x, y float64) (float64, int) {
	form := dropForm(x, y)
	ct := 0
	// exercise 3.3  if z < 0 - it's
	if form < 0. {
		ct = 2
	} else {
		ct = 1
	}
	return form, ct
}

func climbForm(x, y float64) float64 {
	return (math.Sin(x) / x) * (math.Sin(y) / y)
}

func saddleForm(x, y float64) float64 {
	return math.Pow(x, 2)/math.Pow(25, 2) - math.Pow(y, 2)/math.Pow(17, 2)
}

func dropForm(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
