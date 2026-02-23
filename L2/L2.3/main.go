package main

import (
	"fmt"
	"os"
)

// В Go интерфейс — это: 
// • itab  (информация о типе + таблица методов)
// • data  (указатель на конкретное значение)
//
// Если хотя бы одно из этих двух полей не равно nil — интерфейс считается не-nil.
// Именно поэтому "интерфейс, содержащий nil-указатель" != "nil-интерфейс".

func Foo() error {
	var err *os.PathError = nil
	// type - *os.PathError
	// data - nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        //Вывод: <nil>
	fmt.Println(err == nil) // Вывод: false
	// в data хранится nil, но сама переменная не nil
}