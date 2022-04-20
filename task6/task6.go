package main

import (
	"context"
	"fmt"
	"time"
)

func go1(die chan struct{}) {
	for {
		select {
		case <-die:
			fmt.Println("Stop 1")
			return
		}
	}
}

func go2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop 2")
			return
		}
	}
}

func go3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop 3")
			return
		}
	}
}

func go4() {
	t := time.After(time.Second)
	for {
		select {
		case <-t:
			fmt.Println("Stop 4")
			return
		}
	}
}

func main() {
	// Закрытие специального сигнального канала(или отправка в него сигнала)
	die := make(chan struct{})
	go go1(die)
	close(die)

	// Контекст и функция отмены
	ctx, cancel := context.WithCancel(context.Background())
	go go2(ctx)
	cancel()

	// Контекст с ограниченным временем
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))
	go go3(ctx)
	defer cancel()
	
	// Горутина сама завершается через одну секунду
	go go4()
	time.Sleep(2 * time.Second)
}
