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
