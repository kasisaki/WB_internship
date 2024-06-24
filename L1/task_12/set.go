package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

// makeSet создает множество из переданного массива строк
func makeSet(s []string) map[string]struct{} {
	var stub struct{}
	result := make(map[string]struct{})
	for _, str := range s {
		result[str] = stub
	}

	return result
}

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	result := makeSet(s)

	fmt.Printf("Оригинальная последовательность %v\n", s)
	fmt.Print("Получившееся множество: [")

	for k := range result {
		fmt.Print(k + " ")
	}
	fmt.Println("]")
}
