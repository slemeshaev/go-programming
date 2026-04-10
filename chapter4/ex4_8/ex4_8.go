// Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
// using functions like unicode.IsLetter.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letters := 0
	digits := 0
	spaces := 0
	puncts := 0
	others := 0

	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsSpace(r):
			spaces++
		case unicode.IsPunct(r):
			puncts++
		default:
			others++
		}

		utflen[n]++
	}

	fmt.Println("Category\tCount")
	fmt.Printf("Letters\t\t%d\n", letters)
	fmt.Printf("Digits\t\t%d\n", digits)
	fmt.Printf("Spaces\t\t%d\n", spaces)
	fmt.Printf("Punctuation\t%d\n", puncts)
	fmt.Printf("Others\t\t%d\n", others)

	fmt.Print("\nUTF-8 length\tCount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d bytes\t\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\nInvalid UTF-8 characters: %d\n", invalid)
	}
}
