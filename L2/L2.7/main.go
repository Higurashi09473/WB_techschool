package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int) // Небуферизированный канал
	go func() {
		for _, v := range vs {
			c <- v // Записывается одно значение и ждем, когда его считают
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c) // После записи всех чисел канал закрывается
	}()
	return c // возвращаем канал на чтение
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int) // Создается новый канал, который будет объединять каналы
	go func() {
		for {
			select { // select выбирает case с готовым каналом
			case v, ok := <-a:
				if ok { 
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil { // Если оба канала закрыты, то закрываем с
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Print(v) //рандомный вывод чисел например: 12463578 /  12346587
	}
}
