package main

// find sum of two values in an array to target to value two get target value

//array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func two(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			println(arr[left], arr[right])
			return arr[left], arr[right]
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return -1, -1

}
