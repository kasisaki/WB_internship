package main

import (
	"cmp"
	"fmt"
	utils "github.com/kasisaki/WB_internship/task_16/quicksort"
)

// Реализовать бинарный поиск встроенными методами языка.

// binarySearch осуществляет поиск по отсортированному срезу и возвращает индекс искомого значения и true
// если искомый элемент отсутствует в срезе, возвращается -1 и false
// Поиск осуществляется путем сравнение значения в центре(mid) среза с целевым
// если mid == target процесс поиска завершен, если mid < target, то поиск продолжается справа, mid > target -> слева
func binarySearch[S cmp.Ordered](s []S, target S) (int, bool) {
	if len(s) == 0 {
		return -1, false
	}

	if len(s) == 1 {
		if s[0] == target {
			return 0, true
		}
		return -1, false
	}

	mid := len(s) / 2

	if s[mid] == target {
		return mid, true
	}
	if s[mid] > target {
		if localIdx, found := binarySearch(s[:mid], target); found {
			return mid - localIdx, true
		}
		return -1, false
	}
	if s[mid] < target {
		if localIdx, found := binarySearch(s[mid:], target); found {
			return mid + localIdx, true
		}
		return -1, false
	}

	return -1, false
}

func main() {
	a := []int{8, 3, -1, -2, -9, 3, 1, 4, 2, 1, 0, 10, 11, -1, -2}
	utils.Quicksort(a)
	target := 2
	idx, found := binarySearch(a, target)
	fmt.Printf("In %v, target=%d, has index = %d, found: %v\n", a, target, idx, found)
	fmt.Println("------=------")

	target = 3
	idx, found = binarySearch(a, target)
	fmt.Printf("In %v, target=%d, has index = %d, found: %v\n", a, target, idx, found)
	fmt.Println("------=------")

	target = -100500
	idx, found = binarySearch(a, target)
	fmt.Printf("In %v, target=%d, has index = %d, found: %v\n", a, target, idx, found)
	fmt.Println("------=------")

	a = []int{}
	idx, found = binarySearch(a, target)
	fmt.Println("Check if slice is empty")
	fmt.Printf("In %v, target=%d, has index = %d, found: %v\n", a, target, idx, found)
	fmt.Println("------=------")

	a = nil
	idx, found = binarySearch(a, target)
	fmt.Println("Check if slice is nil")
	fmt.Printf("In %v, target=%d, has index = %d, found: %v\n", a, 100, idx, found)
	fmt.Println("------=------")

}
