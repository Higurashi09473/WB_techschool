package main

import "fmt"

type Target interface {
	SomeFunc()
}

type Adaptee struct{}

func (a *Adaptee) SayHello() {
	fmt.Println("Hello from Adaptee!")
}

type Adapter struct {
	*Adaptee 
}

func (a *Adapter) SomeFunc() {
	a.SayHello() 
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{Adaptee: adaptee}
}

func main() {
	client := NewAdapter(&Adaptee{})
	client.SomeFunc() 
}
