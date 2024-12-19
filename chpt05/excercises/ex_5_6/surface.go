/*
Упражнение 5,6. Модифицируйте функцию corner из gopl. io/ch3/surface (раздел 3.2),
чтобы она использовала именованные результаты и инструкцию пустого возврата.
*/
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
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	// fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
	// 	"style=' stroke: grey; fill: white; stroke-width: 0.7' "+
	// 	"width=,%d' height='%d' >", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			file.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	file.WriteString(fmt.Sprintln("</svg>"))
}

func corner(i, j int) (sx float64, sy float64, err error) {
	sx, sy = 0, 0
	err = nil
	// Ищем угловую точку (х,у) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Управжнение 3.1
	if math.IsInf(z, 0) || math.IsNaN(z) {
		err = fmt.Errorf("invalid value")
		return
	}
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0, 0)
	return math.Sin(r)
}
