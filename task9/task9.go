package main

import (
	"fmt"
	"sync"
)

func main() {
	inX := make(chan int)
	outX := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(outX)
		for x := range inX {
			outX <- x * x
		}
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for x := range outX {
			fmt.Println(x)
		}
	}(&wg)

	data := [5]int{1, 2, 3, 4, 5}

	for _, x := range data {
		inX <- x
	}
	close(inX)
	wg.Wait()
}
