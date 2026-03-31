// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice
// that represents a UTF-8-encoded string, in place. Can you do it without
// allocating new memory?

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUTF8(b []byte) []byte {
	runes := make([]rune, 0, utf8.RuneCount(b))
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		runes = append(runes, r)
		i += size
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	result := make([]byte, 0, len(b))
	for _, r := range runes {
		buf := make([]byte, utf8.RuneLen(r))
		utf8.EncodeRune(buf, r)
		result = append(result, buf...)
	}

	return result
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
