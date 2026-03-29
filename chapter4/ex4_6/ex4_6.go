// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(s []byte) []byte {
	if len(s) == 0 {
		return s
	}

	out := 0
	spaceSeq := false

	for in := 0; in < len(s); {
		r, size := utf8.DecodeRune(s[in:])

		if unicode.IsSpace(r) {
			if !spaceSeq {
				s[out] = ' '
				out++
				spaceSeq = true
			}
		} else {
			copy(s[out:out+size], s[in:in+size])
			out += size
			spaceSeq = false
		}
		in += size
	}

	return s[:out]
}

func main() {
	data := []byte("Hello\t\tWorld\n\n\n   Привет\u00A0\u2000мир")
	fmt.Printf("Before: %q\n", string(data))

	result := squashSpaces(data)
	fmt.Printf("After: %q\n", string(result))
}
