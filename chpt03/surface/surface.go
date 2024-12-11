package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // размер канвы в пикселях
	cells         = 100                 // количество ячеек сетки
	xyrange       = 30.0                //диапазон осей -xyrange..+xyrange
	xyscale       = width / 2 / xyrange // пикселей в единце Х или У
	zscale        = height * 0.4        // пикселей в единице Z
	angle         = math.Pi / 6         // Углы осей x, y (=30°))
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin30, cos30

func main() {
	file, err := os.OpenFile("./image.svg", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style=' stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d' >\n", width, height))
	// fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
	// 	"style=' stroke: grey; fill: white; stroke-width: 0.7' "+
	// 	"width=,%d' height='%d' >", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			file.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %gj%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	file.WriteString(fmt.Sprintln("</svg>"))
}

func corner(i, j int) (float64, float64) {
	// Ищем угловую точку (х,у) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0, 0)
	return math.Sin(r)
}
