package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

// Данная реализация будет работать только в версиях 1.20+ из-за TryLock()
// Никакой надобности в TryLock() в этой задаче нет, но просто решил поиграться
func main() {
	var mu sync.Mutex // Mutex позволяет ограничить доступ горутине к участку кода, когда этот участок выполняет другая горутина
	myMap := make(map[string]int)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			ableToLock := mu.TryLock() // Или же можно написать mu.Lock() без цикла и тогда рутина будет ждать пока сможет заблокирвать и затем пойдет дальше по коду
			for !ableToLock {
				fmt.Printf("не получаеца записать  -=%d=-. Повторим!\n", i)
				ableToLock = mu.TryLock()
			}
			fmt.Printf("Прорвались! пишем в мапу число %d\n", i)
			myMap[key] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait() // Подождем, пока рутины сделают свое дело
	fmt.Println("Ща проверим что получилось")
	for k, v := range myMap {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Printf("Хотели записать 10 чисел, а записалось %d\nкто бы сомневался\n", len(myMap))

}
