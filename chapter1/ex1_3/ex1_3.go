// Exercise 1.3. Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func measureTime(name string, fn func(), iterations int) {
	start := time.Now()

	for i := 0; i < iterations; i++ {
		fn()
	}

	elapsed := time.Since(start)
	totalTime := elapsed
	avgTime := time.Duration(int64(elapsed) / int64(iterations))

	fmt.Printf("%s:\n", name)
	fmt.Printf("--Total for %d iterations: %v\n", iterations, totalTime)
	fmt.Printf("--Average per iteration: %v\n", avgTime)
}

func main() {
	measureTime("Echo1", echo1, 100000)
	measureTime("Echo2", echo2, 100000)
	measureTime("Echo3", echo3, 100000)
}

func echo1() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "*"
	}

	_ = s
}

func echo2() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "&"
	}

	_ = s
}

func echo3() {
	result := strings.Join(os.Args[1:], "$")
	_ = result
}
