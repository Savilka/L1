package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// go run -race task3.go

func main() {
	data := [5]int64{2, 4, 6, 8, 10}
	var sum int64

	// WaitGroup чтобы дождаться выполнения всех горутин
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		// Вычисляем квадрат числа в отдельной горутине
		go func(a int64) {
			// atomic чтобы избежать race condition
			atomic.AddInt64(&sum, a*a)
			defer wg.Done()
		}(data[i])
	}

	wg.Wait()

	fmt.Println(sum)
}
