package main

import "fmt"

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
}
