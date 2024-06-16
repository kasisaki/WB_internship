package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Длительность работы программы в секундах берется из аргументов командной строки.
	// Параметр должен предавать как -timeLimit=n, где n - int > 0
	// Значение по умолчанию = 5 секунд
	multiplier := flag.Int("timeLimit", 5, "Program run time limit")
	flag.Parse()

	if *multiplier <= 0 {
		fmt.Printf("Запуск программы имеет смысл только если ей выделено время больше 0 с. Завершаем работу т.к. переданное значение = %d\n", *multiplier)
		return
	}

	duration := time.Duration(*multiplier) * time.Second

	// Создание канала  с типом данных int
	dataChannel := make(chan int)

	// Таймер для завершения программы
	timer := time.NewTimer(duration)

	// Горутина для отправки данных в канал
	go func() {
		for {
			select {
			case <-timer.C:
				close(dataChannel) // Закрытие канала при истечении времени
				return
			default:
				n := rand.Intn(150)
				fmt.Println("Записываем в канал число:", n)
				dataChannel <- n
				time.Sleep(250 * time.Millisecond) // Имитация задержки при отправке данных
			}
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for val := range dataChannel {
			fmt.Println("Получено значение:", val)
		}
		fmt.Println("Чтение из канала завершено")
	}()

	// Ожидание завершения таймера
	<-timer.C
	fmt.Println("Программа завершена")
}
