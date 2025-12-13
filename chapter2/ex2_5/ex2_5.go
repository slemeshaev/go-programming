// Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact, and
// access it performance.

package main

import (
	"fmt"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count
// (number of set bits) of x.
func popCount(x uint64) int {
	count := 0

	for x != 0 {
		x = x & (x - 1)
		count++
	}

	return count
}

func main() {
	fmt.Println(popCount(10))
}
