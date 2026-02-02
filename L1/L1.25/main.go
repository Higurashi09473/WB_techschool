package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	if duration <= 0 {
		return
	}
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	sleep(5 * time.Second)
	fmt.Println("timeout")
}