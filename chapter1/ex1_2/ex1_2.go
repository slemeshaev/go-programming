// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments,
// one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, val := range os.Args {
		fmt.Printf("Index - %d, arg - %s\n", idx, val)
	}
}
