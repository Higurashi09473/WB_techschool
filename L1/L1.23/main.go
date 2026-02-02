package main

import "fmt"

func removeAt(slice []int, i int) ([]int, error) {
	if i >= len(slice) || i < 0 {
		return slice, fmt.Errorf("index out of range")
	}

	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1], nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(removeAt(slice, 4))
}
