// basename убирает компоненты каталога и суффикс типа файла
// a -> a, a.go -> a, a/b/c.go ->c, a/b.c.go -> b.c

package main

import "fmt"

func basename(s string) string {
	// Отбрасываем последний символ '/' и все перед ним.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Сохраняем все до последней точки '.'
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
