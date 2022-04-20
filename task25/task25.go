package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start program")
	var n int
	fmt.Print("How long to sleep in seconds: ")
	fmt.Scan(&n)
	Sleep(time.Second * time.Duration(n))
	fmt.Printf("End program after %d seconds", n)
}
