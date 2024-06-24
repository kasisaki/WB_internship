package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

// AddSquare конкурентно добавляет квадрат указанного числа к сумме, при этом доступ к сумме блокируется пока текущая рутина не закончит работу с этим чсилом
func AddSquare(n int, sum *int) {
	defer wg.Done()

	mu.Lock()
	*sum += n * n
	mu.Unlock()
}

func sumSquares(nums []int) int {
	sum := 0
	for _, n := range nums {
		wg.Add(1)
		go AddSquare(n, &sum)
	}

	wg.Wait() // WaitGroup позволяет подождать пока рутины не закончат свою работу перед тем как вернуть значение суммы
	return sum
}

func main() {
	nums := []int{1, 2, 4, 6, 8, 10, 0}

	fmt.Println(sumSquares(nums))
}
