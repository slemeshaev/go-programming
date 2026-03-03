package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)

	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	s := fmt.Sprintf("x = %b", x) // "x = 1111011"
	fmt.Println(s)

	a, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("Formatting error")
	}

	fmt.Println(a)

	b, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Printf("Formatting error")
	}

	fmt.Println(b)
}
