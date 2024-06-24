package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Counter — это структура, которая включает в себя целочисленный счетчик и мьютекс для управления одновременным доступом.
type Counter struct {
	n  int
	mu sync.Mutex
}

// Increment безопасно увеличивает счетчик на 1.
func (s *Counter) increment() {
	s.mu.Lock()
	s.n++
	s.mu.Unlock()
}

// GetValue безопасно извлекает текущее значение счетчика.
func (s *Counter) GetValue() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.n
}

func main() {
	const goroutineCount = 1500
	var counter Counter
	wg := sync.WaitGroup{}

	wg.Add(goroutineCount)
	for m := 0; m < goroutineCount; m++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				counter.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.GetValue())
}
