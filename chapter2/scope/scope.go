package main

import "fmt"

func f() {}

var g = "g"

// scope
func area() {
	f := "f"

	fmt.Println(f) // "f"; local var f shadows package-level func f
	fmt.Println(g) // "g"; package-level var
	// fmt.Println(h) // compile error: undefined: h

	/*
		if x := f(); x == 0 {
			fmt.Println(x)
		} else if y := g(x); x == y {
			fmt.Println(x, y)
		} else {
			fmt.Println(x, y)
		}
	*/

	// fmt.Println(x, y) // compile error: x and y are not visible here
}

// shadow
func shadow() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}

func main() {
	area()
	shadow()
}
