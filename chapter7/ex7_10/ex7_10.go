// Exercise 7.10: The sort.Interface type can be adapted to other uses.
// Write a function IsPalindrome(s sort.Interface) bool that reports whether
// the sequence s is a palindrome, in othe words, reversing the sequence would not change it.
// Assume that the elements at indeces i and j are equal if !s.Less(i, j) && !s.Less(j, i).

package main

import (
	"fmt"
	"sort"
)

type nums []int

func (s nums) Len() int {
	return len(s)
}

func (s nums) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s nums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - 1 - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	nums1 := nums{1, 2, 2, 1}
	nums2 := nums{1, 2, 2, 3}
	nums3 := nums{1, 2, 2, 3, 4}
	nums4 := nums{}
	nums5 := nums{1}

	fmt.Println(isPalindrome(nums1))
	fmt.Println(isPalindrome(nums2))
	fmt.Println(isPalindrome(nums3))
	fmt.Println(isPalindrome(nums4))
	fmt.Println(isPalindrome(nums5))
}
