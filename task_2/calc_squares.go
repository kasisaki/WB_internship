package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Square рассчитывает квадрат числа и выводит полученный результат
func Square(n int) {
	defer wg.Done()

	fmt.Println(n * n)
}

// CalcSquares расчитывает квадраты чисел из слайса конкурентно путем запуска горутин по расчету и выводу квадратов чисел
func CalcSquares(nums []int) {
	for _, n := range nums {
		wg.Add(1)
		go Square(n)
	}
}

func main() {
	nums := []int{1, 2, 4, 6, 8, 10, 0}

	CalcSquares(nums)
	wg.Wait() // Использование Waitgroup позволяет не завершать main до того, как рутины отработают
}
