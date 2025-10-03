// Exercise 1.1. Change the echo program so that it also outputs os.Args[0],
// the name of the command being executed.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(os.Args)
	fmt.Println(strings.Join(os.Args, "*"))
}
