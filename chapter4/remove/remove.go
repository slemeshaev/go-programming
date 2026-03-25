package main

import "fmt"

func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	r := []int{5, 6, 7, 8, 9}
	fmt.Println(remove1(r, 2)) // [5 6 8 9]

	q := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(q, 2)) // [5 6 9 8]
}
