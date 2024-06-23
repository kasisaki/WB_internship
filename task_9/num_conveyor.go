package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	firstCh := make(chan int)
	secondCh := make(chan int)
	done := make(chan bool)
	nums := []int{0, 2, 3, 1, 5, 100, 12}
	var mu sync.Mutex

	// Рутина будет отправлять числа в канал из массива в случайном порядке
	go func() {
		for {
			select {
			case <-done:
				close(firstCh)
				return
			default:
				time.Sleep(250 * time.Millisecond)
				i := rand.Intn(len(nums) - 1)
				firstCh <- nums[i]
				mu.Lock() // Используем здесь мютекс, чтобы не возникала путаница из-за порядка вывода сообщений
				fmt.Printf("Отправили в первый канал число %d\n", nums[i])
			}
		}
	}()

	// Рутина будет отправлять читать числа из первого канала и отправлять их квадрат в другой
	go func() {
		for n := range firstCh {
			time.Sleep(450 * time.Millisecond)
			secondCh <- n * n
		}
		close(secondCh)
	}()

	go func() {
		for n := range secondCh {
			time.Sleep(600 * time.Millisecond)
			fmt.Printf("Выводим %d\n", n)
			mu.Unlock() // Позволим первой рутине вывести сообщение
		}
	}()

	fmt.Printf("Всего работают %d рутин\n", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(2 * time.Second)
	if runtime.NumGoroutine() == 1 { // main тоже рутина, ее тут не учитываем
		fmt.Println("Все горутины завершили работу")
		return
	}
	fmt.Println("Рутинки работают еще")
}
