// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.

package main

import "fmt"

func removeAdjacentDuplicates(s []string) []string {
	if len(s) <= 1 {
		return s
	}

	write := 1
	for read := 1; read < len(s); read++ {
		if s[read] != s[write-1] {
			s[write] = s[read]
			write++
		}
	}

	return s[:write]
}

func main() {
	a := []string{"a", "a", "b", "b", "b", "c", "d", "d", "a", "a"}
	a = removeAdjacentDuplicates(a)
	fmt.Printf("a: %q\n", a)

	b := []string{"x", "x", "x", "x"}
	b = removeAdjacentDuplicates(b)
	fmt.Printf("b: %q\n", b)

	c := []string{"a", "b", "c", "d"}
	c = removeAdjacentDuplicates(c)
	fmt.Printf("c: %q\n", c)

	d := []string{}
	d = removeAdjacentDuplicates(d)
	fmt.Printf("d: %q\n", d)
}
