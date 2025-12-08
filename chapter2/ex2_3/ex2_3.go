// Exercise 2.3: Rewrite PopCount to use a loop instead of a single
// expression. Compare the perfomance of the two versions. (Section 11.4
// shows how to compare the performance of different implementations
// systematically.)

package main

import (
	"fmt"
	"tempconv/chapter2/popcount"
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
	var sum byte

	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}

	return int(sum)
}

func main() {
	fmt.Println(popcount.PopCount(10))
	fmt.Println(popCount(10))
}
