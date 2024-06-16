package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Чтение количества воркеров из аргументов командной строки
	nWorkers := *flag.Int("workers", 5, "Number of workers")
	flag.Parse()

	// Создание канала для передачи данных
	ch := make(chan any)
	// Создание контекста с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создание канала для сигналов завершения
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Создание WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Горутина для обработки сигналов завершения
	go func() {
		<-sig
		fmt.Println("Получен сигнал завершения")
		cancel()
	}()

	// Горутина для генерации случайных чисел и отправки их в канал
	go func() {
		for {
			select {
			case <-ctx.Done(): // Проверка на отмену контекста
				close(ch)
				return
			default:
				ch <- rand.Intn(1500) // Отправка случайного числа в канал
			}
		}
	}()

	// Запуск воркеров
	work(nWorkers, ch, ctx, &wg)

	// Ожидание завершения всех воркеров
	wg.Wait()
	fmt.Println("Все воркеры завершены")
}

// startWorker обрабатывает данные из канала и завершает работу при получении сигнала отмены контекста
func startWorker(ch chan any, ctx context.Context, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	for {
		time.Sleep(500 * time.Millisecond) // Симуляция работы воркера
		select {
		case <-ctx.Done(): // Проверка на отмену контекста
			fmt.Printf("Воркер %d ушел на пенсию\n", n)
			return
		case data, ok := <-ch: // Чтение данных из канала
			if !ok {
				return
			}
			fmt.Println(data) // Вывод данных
		}
	}
}

// work запускает указанное количество воркеров
func work(workers int, ch chan any, ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i < workers; i++ {
		time.Sleep(185 * time.Millisecond)
		wg.Add(1)                        // Увеличение счетчика воркеров в WaitGroup
		go startWorker(ch, ctx, wg, i+1) // Запуск воркера
	}
}
