package main

import "fmt"

type Human struct {
	Name string
	Age  uint8
}

func (h Human) SayHi() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Action встраивает структуру Human
type Action struct {
	Human
}

func main() {
	a := Action{
		Human: Human{
			Name: "John",
			Age:  30,
		},
	}
	a.SayHi() // Здесь произойдет вызов метода SayHi из Human
}
