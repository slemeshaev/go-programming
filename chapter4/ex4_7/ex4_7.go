// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice
// that represents a UTF-8-encoded string, in place. Can you do it without
// allocating new memory?

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseUTF8(b []byte) []byte {
	if len(b) == 0 {
		return b
	}

	reverseBytes(b)

	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseBytes(b[i : i+size])
		i += size
	}

	return b
}

func main() {
	s := "Hello, 世界"
	data := []byte(s)

	fmt.Printf("Original: %s\n", s)
	fmt.Printf("Bytes: %x\n", data)

	result := reverseUTF8(data)
	fmt.Printf("Reversed bytes: %x\n", result)
	fmt.Printf("Result: %s\n", string(result))
}
