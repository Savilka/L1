package main

import (
	"fmt"
	"strconv"
	"sync"
)

// RWMutex помогает реализовать конкурентную запись и чтение в map

type cMap struct {
	data map[string]string
	mu   sync.RWMutex
}

func (m *cMap) Add(key, value string) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}

func (m *cMap) Get(key string) (string, bool) {
	m.mu.RLock()
	value, ok := m.data[key]
	m.mu.RUnlock()
	return value, ok
}

func main() {
	m := cMap{
		data: make(map[string]string),
		mu:   sync.RWMutex{},
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			m.Add("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
			defer wg.Done()
		}(&wg, i)
	}

	wg.Wait()
	for i := 0; i < len(m.data); i++ {
		value, _ := m.Get("key" + strconv.Itoa(i))
		fmt.Println(value)
	}

}
