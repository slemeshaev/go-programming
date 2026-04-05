package main

import (
	"fmt"
	"sort"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	age, ok := ages["bob"]
	if !ok {
		fmt.Println("bob is not key in this map")
	} else {
		fmt.Printf("%s is a key in this map %d\n", "bob", age)
	}

	if age, ok := ages["alice"]; !ok {
		fmt.Println("alice is not key in this map")
	} else {
		fmt.Printf("%s is a key in this map %d\n", "alice", age)
	}

	fmt.Println(equal(map[string]int{"A": 42}, map[string]int{"B": 42}))
}
