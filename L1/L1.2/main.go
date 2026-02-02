package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	wg.Add(5)
	for index := range len(array) {
		go func(i int) {
			defer wg.Done()
			array[i] = array[i] * array[i]
		}(index)
	}
	wg.Wait()
	fmt.Println(array)
}

//	func main() {
//		array := [5]int{2, 4, 6, 8, 10}
//		wg := &sync.WaitGroup{}
//		wg.Add(5)
//		for index := range len(array) {
//			go func(i int){
//				defer wg.Done()
//				var num int = array[i]
//				fmt.Println(num * num)
//			}(index)
//		}
//		wg.Wait()
//	}

/*
Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/
