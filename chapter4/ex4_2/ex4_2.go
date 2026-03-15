// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	sha384flag := flag.Bool("sha384", false, "Use SHA384 hash instead of SHA256")
	sha512flag := flag.Bool("sha512", false, "Use SHA512 hash instead of SHA256")

	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	switch {
	case *sha384flag:
		hash := sha512.Sum384(data)
		fmt.Printf("%x\n", hash)
	case *sha512flag:
		hash := sha512.Sum512(data)
		fmt.Printf("%x\n", hash)
	default:
		hash := sha256.Sum256(data)
		fmt.Printf("%x\n", hash)
	}
}
