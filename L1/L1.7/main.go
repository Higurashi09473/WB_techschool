package main

import (
	"fmt"
	"sync"
)

func main() {
	set := make(map[int]string)
	mt := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mt.Lock()
			set[i] = fmt.Sprint("data ", i)
			mt.Unlock()
		}(i)
	}

	wg.Wait()
	for j := 0; j < 100; j++ {
		fmt.Println(set[j])
	}
}
