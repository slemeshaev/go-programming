package main

import "fmt"

func main() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // 0
	fmt.Println(q[2]) // 3
}
