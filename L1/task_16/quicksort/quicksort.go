package utils

import (
	"cmp"
	"math/rand"
)

func Quicksort[S cmp.Ordered](s []S) {

	l := len(s)
	// Если срез пустой или в нем только один элемент, то нечего сортировать
	if l < 2 {
		return //
	}

	// Если массив состоит из двух значений, то смысла дальше идти нет и просто сортируем их на месте и выходим
	if l == 2 {
		if s[0] > s[1] {
			s[0], s[1] = s[1], s[0]
		}
		return // точка выхода
	}
	// Выбираем опорный элемент, левый и правый индексы
	// опорный элемент выбирается случайным образом для снижения вероятности возможности злонамеренного подбора наиболее неподходящих(сложных) входных данных
	pivotIdx := rand.Intn(len(s) - 1)
	leftIdx, rightIdx := 0, l-2

	// перемещаем в конец опорное значение
	s[pivotIdx], s[l-1] = s[l-1], s[pivotIdx]
	pivot := s[l-1]
	//fmt.Printf("Arr: %v, pivot[%d]\n", s, pivotIdx)

	for leftIdx < rightIdx {
		// Двигаем левый указатель вправо, пока не найдем элемент, больший или равный опорному.
		for s[leftIdx] < pivot && leftIdx < rightIdx {
			leftIdx++
		}
		// Двигаем правый указатель влево, пока не найдем элемент, меньший или равный опорному.
		for s[rightIdx] > pivot && leftIdx < rightIdx {
			rightIdx--
		}
		// сравниваем значение под левым и правым индексом и меняем местами, если левый больше правого
		if s[leftIdx] > s[rightIdx] {
			s[leftIdx], s[rightIdx] = s[rightIdx], s[leftIdx]
		}
	}
	// Определяем новое положение опорного значения: если опорное меньше значения под правым индексом, то они поменяются местами
	if s[rightIdx] > pivot {
		s[rightIdx], s[pivotIdx] = s[pivotIdx], s[rightIdx]
	}

	Quicksort(s[:leftIdx+1])
	Quicksort(s[rightIdx:])

	return
}
