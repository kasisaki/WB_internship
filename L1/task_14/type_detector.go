package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}.

var Reset = "\033[0m"
var Green = "\033[32m"

func getType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "unknown type"
	}
}

func main() {
	a := 3
	b := "naughty string"
	d := make(chan int)

	// Определяем тип с помощью reflect.TypeOf(i any)
	fmt.Println("Определяем тип с помощью reflect.TypeOf(v any)")
	func(a interface{}, b interface{}, c interface{}, d interface{}) {
		fmt.Printf("a has type %s%s%s\n", Green, reflect.TypeOf(a), Reset)
		fmt.Printf("b has type %s%s%s\n", Green, reflect.TypeOf(b), Reset)
		fmt.Printf("c has type %s%s%s\n", Green, reflect.TypeOf(c), Reset)
		fmt.Printf("d has type %s%s%s\n", Green, reflect.TypeOf(d), Reset)
	}(a, b, true, d)

	// Определяем тип с помощью %T в fmt.Printf
	fmt.Println("Определяем тип с помощью %T в fmt.Printf")
	func(a interface{}, b interface{}, c interface{}, d interface{}) {
		fmt.Printf("a has type %s%T%s\n", Green, a, Reset)
		fmt.Printf("b has type %s%T%s\n", Green, b, Reset)
		fmt.Printf("c has type %s%T%s\n", Green, c, Reset)
		fmt.Printf("d has type %s%T%s\n", Green, d, Reset)
	}(a, b, true, d)

	// Определяем тип с помощью a.(type) в switch
	fmt.Println("Определяем тип с помощью v.(type) в switch")
	fmt.Printf("a has type %s%s%s\n", Green, getType(a), Reset)
	fmt.Printf("b has type %s%s%s\n", Green, getType(b), Reset)
	fmt.Printf("c has type %s%s%s\n", Green, getType(true), Reset)
	fmt.Printf("d has type %s%s%s\n", Green, getType(d), Reset)
}
