// basename removes directory components and a .suffix
// a -> a, a.go -> a, a/b/c.go ->c, a/b.c.go -> b.c

package main

import "fmt"

func basename(s string) string {
	// discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
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
