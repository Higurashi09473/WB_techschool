package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)


func main() {
	exit()
	exitStopChan()
	exitCtx()
	exitRuntime()
	exitTime()
	exitChan()
}

// пример остановки горутины по условию
func exit() {
	var stop bool
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for !stop {
			fmt.Println("running...")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("goroutine stopped")
	}()
	time.Sleep(time.Second)
	stop = true
	wg.Wait()
	fmt.Println("exit")
}

// пример остановки горутины через канал
func exitStopChan() {
	stopCh := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				fmt.Println("Exit goroutine")
				return
			default:
				fmt.Println("Running goroutine")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	stopCh <- struct{}{}
	wg.Wait()
	fmt.Println("exit")
}

// пример остановки горутины через контекст
func exitCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Exit goroutine")
				return
			default:
				fmt.Println("Running goroutine")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}

// пример остановки горутины с помощью runtime.Goexit
func exitRuntime() {
	go func(){
		for i := 0; i <= 5; i++{
			fmt.Println("Gopher works")
		}
		fmt.Println("Gopher left")
		runtime.Goexit()
	}()
}


// пример остановки через время
func exitTime() {
	timeCh := time.After(time.Second * time.Duration(3))
	
	go func()  {
		for i := 1; ; i++ {
			select {
			case <- timeCh:
				fmt.Println("Gopher left")
				return
			default:
				fmt.Println("Gopher works")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}
	
// пример остановки горутины с использованием канала для получения данных
func exitChan() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch{
			fmt.Println(value)
		}
		}()
	
	for i := 0; i < 5; i++{
		ch <- "Gopher work"
	}

	close(ch)
	wg.Wait()
	fmt.Println("exit")
}

/*
Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления, через контекст,
прекращение работы runtime.Goexit() и др.

Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/