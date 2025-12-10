// Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
// bit positions, testing the rightmost bit each time. Compare its performance to the tablelookup version.

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
		count += int(x & 1)
		x >>= 1
	}

	return count
}

func main() {
	fmt.Println(popCount(10))
}
