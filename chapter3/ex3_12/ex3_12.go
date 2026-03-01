// Exercise 3.11: Write a function that reports whether two strings
// are anagrams of each other, that is, they contain the same letters in
// a different order.

package main

import (
	"fmt"
	"strings"
)

func areAnagrams(s1, s2 string) bool {
	s1 = strings.ReplaceAll(strings.ToLower(s1), " ", "")
	s2 = strings.ReplaceAll(strings.ToLower(s2), " ", "")

	if len(s1) != len(s2) {
		return false
	}

	count := make(map[rune]int)
	for _, ch := range s1 {
		count[ch]++
	}

	for _, ch := range s2 {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		a, b string
	}{
		{"listen", "silent"},
		{"triangle", "integral"},
		{"debit card", "bad credit"},
		{"hello", "world"},
		{"rat", "car"},
		{"это", "тоэ"},
	}

	for _, tc := range testCases {
		fmt.Printf("%q и %q: %v\n", tc.a, tc.b, areAnagrams(tc.a, tc.b))
	}
}
