// Surface вычисляет SVG-представление трехмерного графика функции
package main

import "math"

const (
	width, height = 600, 300            // размер канвы в пикселях
	cells         = 100                 // количество ячеек сетки
	xyrange       = 30                  // диапазон (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // пикселей в единице x или y
	zscale        = height * 0.4        // пикселей в единице z
	angle         = math.Pi / 6         // углы осей x, y (=30*)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30*), cos(30*)

func main() {
	//
}
