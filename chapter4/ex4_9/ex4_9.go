// Exercise 4.9: Write a program wordfreq to report the frequency of each word
// in an input text file. Call input.Split(bufio.ScanWords) before the first call
// to Scan to break the input into words instead of lines.

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wordfreq <filename>\n")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
}
