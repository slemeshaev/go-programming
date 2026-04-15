package main

import "fmt"

type Point struct {
	X, Y int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)

	fmt.Println(Scale(p, 5))

	pp := &Point{3, 4}
	fmt.Println(pp)

	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)
}
