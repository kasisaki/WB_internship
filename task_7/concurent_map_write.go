package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.
// Данная реализация будет работать только в версиях 1.20+ из-за TryLock()
func main() {
	var mu sync.Mutex // Mutex позволяет ограничить доступ горутине к участку кода, когда этот участок выполняет другая горутина
	myMap := make(map[string]int)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			ableToLock := mu.TryLock()
			for !ableToLock {
				if !ableToLock {
					fmt.Printf("не получаеца записать  -=%d=-. Повторим!\n", i)
					ableToLock = mu.TryLock()
					continue
				}

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
