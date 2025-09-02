package main

import (
	"fmt"
)

func setBit(num int64, pos int64, bit int) int64 {

	switch bit {
		case 1:
			num |= 1 << pos - 1
		case 0:
			num &^= 1 << pos - 1
	}
	return num
}

func main() {
	fmt.Println(setBit(5, 1, 0))
}