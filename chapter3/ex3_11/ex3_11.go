// Exercise 3.11: Enhance comma so that it deals
// correctly with floating-point numbers and an optional sign.

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaInteger(s string) string {
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

func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	var buf bytes.Buffer
	start := 0

	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		start = 1
	}

	dotIndex := strings.Index(s[start:], ".")
	if dotIndex == -1 {
		buf.WriteString(commaInteger(s[start:]))
	} else {
		dotIndex += start
		buf.WriteString(commaInteger(s[start:dotIndex]))
		buf.WriteString(s[dotIndex:])
	}

	return buf.String()
}

func main() {
	testCases := []string{
		"123",
		"12345",
		"1234567",
		"+1234567",
		"-1234567",
		"12345.6789",
		"-12345.6789",
		"+12345.6789",
		"1234567890.12345",
		"-1234567890.12345",
		"0.123",
		"-0.123",
		".123",
		"-.123",
	}

	for _, tc := range testCases {
		fmt.Printf("%-20s -> %s\n", tc, comma(tc))
	}
}
