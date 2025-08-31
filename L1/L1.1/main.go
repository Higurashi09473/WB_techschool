package main

import "fmt"

type Human struct {
	Name string
	Age int
}

func (h *Human) Speak() {
    fmt.Printf("%s говорит: Привет, мне %d лет!\n", h.Name, h.Age)
}

func (h *Human) Walk() {
    fmt.Println(h.Name, "идёт пешком.")
}

func (a *Action) Run() {
	a.Energy--
    fmt.Println(a.Name, "бежит")
	fmt.Println(a.Energy)
}

// Дочерняя структура Action, встраивающая Human (embedding)
type Action struct {
    Human      // Встраивание
    Energy int
}

func main() {
	higu := &Action{
		Human: Human{
			Name: "Higu",
			Age: 19,
		},
		Energy: 20,
	}
	higu.Walk()
	higu.Run()
}



/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/


/* композиция (composition) / встраивание структур (embedding) */