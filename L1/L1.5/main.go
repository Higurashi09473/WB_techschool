package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	var N int
	fmt.Scan(&N)
	timeCh := time.After(time.Second * time.Duration(N))

	go func()  {
		for i := 1; ; i++ {
			select {
				case <- timeCh:
					close(ch)
					return
				default:
					ch <- i
			}
		}
	}()

	for value := range ch {
		fmt.Println(value)
	}
}
