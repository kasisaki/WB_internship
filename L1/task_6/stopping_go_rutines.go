package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//перечисли все возможные способы остановки выполнения горутины.

/*
	Часто говорят, что таких способа 3, но:

	1) завершение main функции и main горутины;

	2) прослушивание всеми горутинами channel, при закрытии channel отправляется значение по умолчанию всем слушателям,
		при получении сигнала все горутины делают return;

	3) завязать все горутины на переданный в них context.

	4) Возможно не самый оптимальный, но возможный --> рутина читает глобальную переменную и при определенном ее значении завершает свою работу

	5) Я такое еще не видел и может выглядеть глупо, но можно реализовать завершение рутины при неудачной попытке заблокировать при помощи мютекса
*/

var stopGlobalist = false
var muGlobal sync.Mutex

func workerChan(id int, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println("+++++++++=+++++++++")
			fmt.Printf("WorkerChan %d stopping\n", id)
			fmt.Println("+++++++++=+++++++++")
			return
		default:
			fmt.Printf("WorkerChan %d working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerCtx(id int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("---------=---------")
			fmt.Printf("WorkerCtx %d stopping\n", id)
			fmt.Println("---------=---------")
			return
		default:
			fmt.Printf("WorkerCtx %d working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerLocker(id int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if mu.TryLock() {
			mu.Unlock()
			fmt.Printf("workerLocker %d acquired the lock and is working\n", id)
			time.Sleep(500 * time.Millisecond) // Имитируем работу
		} else {
			fmt.Println("---------=---------")
			fmt.Printf("workerLocker %d could not acquire the lock and is stopping\n", id)
			fmt.Println("---------=---------")
			return
		}
	}
}

func workerGlobalist(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		muGlobal.Lock()
		if stopGlobalist {
			fmt.Println("---------=---------")
			fmt.Printf("WorkerGlobalist %d stopping as stopGlobalist is %v\n", id, stopGlobalist)
			fmt.Println("---------=---------")
			muGlobal.Unlock()
			return
		}
		muGlobal.Unlock()
		fmt.Printf("WorkerGlobalist %d working\n", id)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var mu sync.Mutex
	stopCh := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Второй вариант
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerChan(i, stopCh, &wg)
	}

	//// третий вариант
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerCtx(i, ctx, &wg)
	}

	// четвертый вариант
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerGlobalist(i, &wg)
	}

	// пятый вариант
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerLocker(i, &mu, &wg)
	}

	time.Sleep(1 * time.Second)
	close(stopCh)

	time.Sleep(2 * time.Second)
	cancel()

	// дадим еще немного поработать воркерам с локом
	time.Sleep(time.Second)

	mu.Lock() // Завершим воркеры с локом
	time.Sleep(2 * time.Second)
	mu.Unlock()

	muGlobal.Lock()
	stopGlobalist = true // завершаем воркеры, которые читают глобальную переменную, мютекс для избежания гонки данных
	muGlobal.Unlock()

	// Wait a little to see the output from goroutines before main exits
	fmt.Println("----Waiting All workers to stop----")
	wg.Wait()
	fmt.Println("All workers stopped")
}
