// basename removes directory components and a .suffix
// a -> a, a.go -> a, a/b/c.go ->c, a/b.c.go -> b.c

package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1, if "/" not found
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func main() {
	str1 := "a"
	str2 := "a.go"
	str3 := "a/b/c.go"
	str4 := "a/b.c.go"

	fmt.Println(basename(str1))
	fmt.Println(basename(str2))
	fmt.Println(basename(str3))
	fmt.Println(basename(str4))
}
