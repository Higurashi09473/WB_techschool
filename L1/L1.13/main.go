package main

import "fmt"

func main() {
	fmt.Println(swap(3, 7))
}
func swap(a,b int) (int,int) {
	a = a + b
	b = a - b // Теперь b = a
	a = a - b // Теперь a = b
	return a,b
}