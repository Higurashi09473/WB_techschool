package main

import (
	"fmt"
)

func main() {
	outch := make(chan int)
	inch := make(chan int)

	nums := []int{1, 3, 5, 7, 9, 11}
	go func() {
		for _, value := range nums {
			inch <- value
		}
		close(inch)
	}()

	go func() {
		for value := range inch {
			outch <- value * 2
		}
		close(outch)
	}()

	for v := range outch {
		fmt.Println(v)
	}
	fmt.Println("Exit")
}
