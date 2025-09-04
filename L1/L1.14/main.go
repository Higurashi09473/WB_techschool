package main

import "fmt"

func main() {
	seeType(1)
	seeType("string")
	seeType(true)
	seeType(make(chan int))
	seeType(make(chan string))
	seeType(make(chan bool))
}

func seeType[T any](v T){
	switch any(v).(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	}
}