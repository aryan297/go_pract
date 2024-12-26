package main

import (
	"fmt"
	"sort"
)

// sum of three values in an array that sum to zero

func sumZero(arr []int) []int {

	sort.Ints(arr)
	found := false

	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == 0 {
				fmt.Println("values found", sum)
				found = true
				left++
				right--
				return []int{arr[i], arr[left], arr[right]}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	if !found {
		fmt.Println("No such values found")
	}

	return []int{}

}
