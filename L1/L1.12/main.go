package main

import "fmt"


func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	result := []string{}

	set := make(map[string]struct{})

	for _, word := range words{
		set[word] = struct{}{}
	}

	for key := range set{
		result = append(result, key)
	}

	fmt.Println(result)
}