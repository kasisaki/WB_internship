package main

import "fmt"

/*
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
*/

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float32)

	for _, n := range temps {
		bin := 10 * int(n/10) // Определяем группу, к которой относится температура

		// Если необходимая группа уже есть в мапе, то добавляем текущее значение к слайсу
		if arr, ok := result[bin]; ok {
			result[bin] = append(arr, n)
			continue
		}
		// иначе добавляем группу и соответсвующую температуру в мапу
		result[bin] = []float32{n}
	}

	fmt.Println(result)
}
