package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// Программа некорректно работает на windows, т.к. windows не поддерживает unix сигналы
// Воркеры прекращают работу при закрытии сигнального канала die
// Ещё есть возможность остановить все горутины с помощью контекста, но я решил, что для такой задачи использвать более
// мощный инструмент излишне

func worker(wg *sync.WaitGroup, die chan struct{}, id int, in chan []byte) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-in:
			if !ok {
				return
			}
			_, err := fmt.Fprintf(os.Stdout, "worker %d print %s \n", id, data)
			if err != nil {
				return
			}
		case <-die:
			_, err := fmt.Fprintf(os.Stdout, "worker %d stoped\n", id)
			if err != nil {
				return
			}
			return
		}
	}

}

func main() {

	wg := sync.WaitGroup{}
	in := make(chan []byte)
	die := make(chan struct{})
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(&wg, die, i, in)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		s := <-sigChan
		fmt.Printf("%s\n", s.String())
		close(die)
		wg.Wait()
		close(in)
		fmt.Println("Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in <- scanner.Bytes()
	}
}
