package main

import "fmt"



func main() {
	arr := []int{1, 4, 6, 9, 12, 16, 17, 19, 20, 400, 456, 500, 1000}
	fmt.Println(binarySearch(arr, 1000))
}

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if arr[mid] == target{
			return mid
		}

		if arr[mid] > target{
			right = mid
		}else {
			left = mid
		}
	}
	return -1
}