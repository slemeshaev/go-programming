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

func corner(i, j int) (float64, float64) {
	// ищем угловую точку (x, y) ячейки (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// вычисялем высоту поверхности z
	z := f(x, y)
	// изметрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // расстояние от (0, 0)
	return math.Sin(r) / r
}
