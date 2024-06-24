package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.
// Предположил, что имеется в виду поиск пересечения двух множеств

func fillMap(s []int, m map[int]bool) {
	for _, n := range s {
		m[n] = true
	}
}

func findIntersection(s []int, m map[int]bool) []int {
	var result []int
	for _, n := range s {
		if m[n] {
			result = append(result, n)
		}
	}
	return result
}

func intersection(s1 []int, s2 []int) []int {
	setMap := make(map[int]bool)

	// Данный участок кода может сохранить немного памяти
	if len(s1) > len(s2) {
		fillMap(s2, setMap)
		return findIntersection(s1, setMap)
	}
	fillMap(s2, setMap)
	return findIntersection(s1, setMap)
}

func main() {
	s1 := []int{0, 2, 3, 1, 5, 100, 12}
	s2 := []int{70, 2, 3, 4, 6, 10, 121}

	fmt.Println(intersection(s1, s2)) // Ожидаем на выходе [2, 3]
}
