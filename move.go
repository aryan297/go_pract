package main

func moveZero(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] == 0 {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		} else {
			left++

		}
	}

}
