package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendsValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func main() {
	// Example 1: Unsorted array
	numbers := []int{5, 3, 8, 1, 4, 9, 2, 7, 6}

	fmt.Println("Before sorting:", numbers)
	Sort(numbers)
	fmt.Println("After sorting:", numbers)

	// Example 2: Already sorted
	numbers2 := []int{1, 2, 3, 4, 5}
	fmt.Println("\nBefore sorting:", numbers2)
	Sort(numbers2)
	fmt.Println("After sorting:", numbers2)

	// Example 3: Reverse order
	numbers3 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("\nBefore sorting:", numbers3)
	Sort(numbers3)
	fmt.Println("After sorting:", numbers3)

	// Example 4: with duplicates
	numbers4 := []int{5, 2, 5, 1, 5, 3, 5, 4}
	fmt.Println("\nBefore sorting:", numbers4)
	Sort(numbers4)
	fmt.Println("After sorting:", numbers4)
}
