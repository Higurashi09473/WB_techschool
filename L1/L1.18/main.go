package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32

	wg := sync.WaitGroup{}

	wg.Add(5)
	for range 5 {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}