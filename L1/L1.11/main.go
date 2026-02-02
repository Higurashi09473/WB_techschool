package main

import "fmt"


func main() {
	slice1 := []int{1,2,3}
	slice2 := []int{2,3,4}

	result := []int{}

	set := make(map[int]struct{})

	for _, value := range slice1 {
		set[value] = struct{}{}
	}

	for _, value := range slice2{
		if _, ok := set[value]; ok {
			result = append(result, value)
		}
	}

	fmt.Println(result)
}