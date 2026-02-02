package main

import "fmt"

func reverse(arr []rune) {
	l := 0
	r := len(arr) - 1
	for l < r {
		arr[r], arr[l] = arr[l], arr[r]
		l++
		r--
	}
}

func main() {
	var n string
	fmt.Scan(&n)
	str := []rune(n)
	reverse(str)
	fmt.Println(string(str))
}