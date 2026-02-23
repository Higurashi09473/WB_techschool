package main

import "fmt"

func test() (x int) { 
	// Именованное возвращение
	// defer-функции могут изменять эту переменную даже после выполнения return
	defer func() { 
		x++
	}()
	x = 1
	return // Так как return возвращает ИМЕННО x, то defer сможет внести изменения в результат
}

func anotherTest() int { 
	// Безымянное возвращение
	// при вызове return создаётся временная копия того значения, которое мы возвращаем
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x // возращаем копию 
	// defer выполняется уже после того, как копия возвращаемого значения создана
}

func main() {
	fmt.Println(test())        // 2
	fmt.Println(anotherTest()) // 1
}
