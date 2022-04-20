package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// go run task18.go -race
// Счетчик в конкурентной среде можно реализовать через атомарные операции
// Также можно использовать mutex, но это более затратная реализация
// Benchmark1-8    277698967                4.356 ns/op
// Benchmark2-8    127808389                9.285 ns/op

type Counter struct {
	count int64
}

func (c *Counter) add() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) get() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	var c Counter
	wg := sync.WaitGroup{}
	for i := 0; i < 55; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			c.add()
		}(&wg)
	}
	wg.Wait()
	fmt.Println(c.get())
}

type CounterMutex struct {
	count int64
	mu    sync.Mutex
}

func (c *CounterMutex) add() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *CounterMutex) get() int64 {
	c.mu.Lock()
	a := c.count
	c.mu.Unlock()
	return a
}
