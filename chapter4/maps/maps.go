package main

import "fmt"

func main() {
	marks := make(map[string]int)
	fmt.Println(marks)

	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Println(numbers)

	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	ages["alice"] = 32
	fmt.Println(ages["alice"]) // 32

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
