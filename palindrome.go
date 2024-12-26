package main

import "fmt"

func plaindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			fmt.Println("Not a palindrome")
			return false
		}
		left++
		right--

	}
	return true
}
