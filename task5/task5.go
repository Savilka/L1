package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Для остановки программы использовалась функция After, которая после заданного времени
// отправляет в канал текущее время, что позволяет завершить программу по таймеру

func main() {
	in := make(chan int)
	die := make(chan bool)

	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, in chan int, die chan bool) {
		defer wg.Done()
		for {
			select {
			case val := <-in:
				fmt.Println(val)
			case <-die:
				fmt.Println("closed")
				return
			}
		}
	}(&wg, in, die)

	var i int
	timer := time.After(time.Second * time.Duration(n))
	for {
		select {
		case <-timer:
			close(die)
			wg.Wait()
			os.Exit(0)
		default:
			in <- i
			i++
			// Sleep для человекочитаемого вывода в консоль
			time.Sleep(100 * time.Millisecond)
		}
	}
}
