package main

import (
	"fmt"
	"sort"
)

func main() {
	var names []string

	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
