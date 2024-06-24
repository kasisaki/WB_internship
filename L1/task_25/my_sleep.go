package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep(duration time.Duration) {
	var mu sync.Mutex
	timer := time.NewTimer(duration)

	for {
		select {
		case <-timer.C:
			mu.Unlock()
			return
		default:
			mu.TryLock()
		}
	}
}

func sleep2(duration time.Duration) {
	var wg sync.WaitGroup
	timer := time.NewTimer(duration)
	wg.Add(1)
	go func() {
		select {
		case <-timer.C:
			wg.Done()
		}
	}()
	wg.Wait()
}

func main() {
	fmt.Println("Время спать")
	sleep(3 * time.Second)
	fmt.Println("Хорошо поспали")

	fmt.Println("Еще раз")
	sleep2(3 * time.Second)
	fmt.Println("Хорошо поспали")
}
