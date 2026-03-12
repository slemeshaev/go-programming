package main

import "fmt"

func main() {
	var x [3]int             // array of 3 integers
	fmt.Println(x[0])        // print the first element
	fmt.Println(x[len(x)-1]) // print the last element, x[2]

	// Print the indices and elements
	for i, v := range x {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only
	for _, v := range x {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // 0
	fmt.Println(q[2]) // 3

	h := [...]int{1, 2, 3}
	fmt.Printf("%T\n", h) // "[3]int"

	k := [3]int{1, 2, 3}
	fmt.Println(k)
	// compile error: cannot assign [4]int to [3]int
	// k = [4]int{1, 2, 3, 4}

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RUR
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RUR: "₽"}
	fmt.Println(RUR, symbol[RUR])

	m := [...]int{99: -1}
	fmt.Println(m)

	// comparable arrays
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}

	// compile error: cannot compare [2]int == [3]int
	// fmt.Println(a == d)
}
