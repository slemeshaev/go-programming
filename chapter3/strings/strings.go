package main

import (
	"fmt"
	"unicode/utf8"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {
	s := "hello, world"

	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

	c := s[len(s)-1] // panic: index out of range
	fmt.Println(c)

	fmt.Println(s[0:5]) // "hello"

	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	// The + operator makes a new string by concatenating two strings
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	foot := "left foot"
	t := foot
	foot += ", right foot"

	fmt.Println(foot) // "left foot, right foot"
	fmt.Println(t)    // "left foot"

	// foot[0] = 'L' // compile error: cannot assign to s[0]

	// example of use
	fmt.Println(HasPrefix(s, "hel")) // true
	fmt.Println(HasPrefix(s, "el"))  // false

	fmt.Println(HasSuffix(s, "rld")) // true
	fmt.Println(HasSuffix(s, "rl"))  // false

	fmt.Println(Contains(s, "o,")) // true
	fmt.Println(Contains(s, "xf")) // false

	str := "Hello, 世界"
	fmt.Println(len(str))                    // "13"
	fmt.Println(utf8.RuneCountInString(str)) // "9"

	// the clumsy solution
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for range str {
		n++
	}

	// "program" in Japanese katakana
	js := "プログラム"
	fmt.Printf("% x\n", js) // e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0

	jr := []rune(js)
	fmt.Printf("%x\n", jr) // [30d7 30ed 30b0 30e9 30e0]

	fmt.Println(string(jr))      // プログラム
	fmt.Println(string(65))      // A
	fmt.Println(string(0x4eac))  // 京
	fmt.Println(string(1234567)) // �
}
