package main

import (
	"fmt"
	"strings"
)

func reverse(arr []string) {
	l := 0
	r := len(arr) - 1
	for l < r {
		arr[r], arr[l] = arr[l], arr[r]
		l++
		r--
	}
}

func main() {
	strs := strings.Fields("snow dog sun")
	reverse(strs)
	fmt.Println(strs)
}