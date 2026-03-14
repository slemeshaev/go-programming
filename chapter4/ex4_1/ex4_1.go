// Exercise 4.1: Write a function that counts the number of bits
// that are different in two SHA256 hashes. (See PopCount from Section 2.6.2)

package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func SHA256DiffBits(h1, h2 [32]byte) int {
	diff := 0
	for i := 0; i < 32; i++ {
		diff += PopCount(uint64(h1[i] ^ h2[i]))
	}
	return diff
}

func main() {
	h1 := sha256.Sum256([]byte("x"))
	h2 := sha256.Sum256([]byte("X"))

	fmt.Printf("SHA256 of 'x': %x\n", h1)
	fmt.Printf("SHA256 of 'X': %x\n", h2)
	fmt.Printf("Different bits: %d\n", SHA256DiffBits(h1, h2))
}
