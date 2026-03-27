// Exercise 4.4: Write a version of rotate that operates in a single pass.

package main

import "fmt"

func rotateLeft(s []int, n int) {
	if len(s) == 0 {
		return
	}

	n = n % len(s)
	if n == 0 {
		return
	}

	temp := make([]int, n)

	for i := 0; i < n; i++ {
		temp[i] = s[i]
	}

	for i := n; i < len(s); i++ {
		s[i-n] = s[i]
	}

	for i := 0; i < n; i++ {
		s[len(s)-n+i] = temp[i]
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotateLeft(a, 2)
	fmt.Println(a) // [2 3 4 5 0 1]
}
