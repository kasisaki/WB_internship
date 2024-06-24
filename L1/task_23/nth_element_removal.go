package main

import "fmt"

// Удалить i-ый элемент из слайса.
func main() {
	nums := []int{1, 2, 0, 4, 5}

	fmt.Printf("Изначальный слайс : %v\n", nums) // Вывод: [1 2 3 4 5]

	// Удаляем 3-й элемент (индекс 2)
	i := 2
	nums = removeElement(nums, i)

	fmt.Printf("После удаления 3го значения : %v\n", nums) // Вывод: [1 2 4 5]

	fmt.Println("---=Еще один способ:")

	nums2 := []int{1, 2, 3, 3, 4, 5}
	fmt.Printf("Изначальный слайс : %v\n", nums2) // Вывод: [1 2 3 3 4 5]

	nums2 = yetAnotherRemover(nums2, 2)

	fmt.Printf("После удаления 3го значения: %v\n", nums2) // Вывод: [1 2 3 4 5]
	yetAnotherRemover([]int{1, 2}, 500)
}

func removeElement(slice []int, index int) []int {
	// Проверяем, чтобы передаваемый индекс был в пределах среза
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Создаем новый слайс, исключая i-ый элемент
	return append(slice[:index], slice[index+1:]...)
}

func yetAnotherRemover(slice []int, targetIdx int) []int {
	// Проверяем, чтобы передаваемый индекс был в пределах среза
	// превышение индекса для данной реализации не страшен, но зачем лишняя работа
	if targetIdx < 0 || targetIdx >= len(slice) {
		return slice
	}

	for i, n := range slice {
		if i <= targetIdx {
			continue
		}
		slice[i-1] = n
	}
	return slice[:len(slice)-1]
}
