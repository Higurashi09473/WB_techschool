package main

import (
	"fmt"
	"wb/unpack"
)

func main() {
	s, err := unpack.Unpack("g23")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}