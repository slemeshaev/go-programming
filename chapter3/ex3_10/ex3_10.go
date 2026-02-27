// Exercise 3.10: Write a non-recursive version of comma,
// using bytes.Buffer instead of string concatenation.

package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	remainder := n % 3
	if remainder == 0 {
		remainder = 3
	}

	buf.WriteString(s[:remainder])

	for i := remainder; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

func main() {
	str := "12345"
	fmt.Println(comma(str))
}
