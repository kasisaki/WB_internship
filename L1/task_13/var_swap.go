package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	a := 2
	b := 9
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Println("Swap in place")
	fmt.Printf("a = %d, b = %d\n", a, b)

	swap(&a, &b)
	fmt.Println("Swap back by function")
	fmt.Printf("a = %d, b = %d\n", a, b)

}
