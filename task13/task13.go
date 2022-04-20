package main

import "fmt"

func main() {
	a := 2
	b := 7
	fmt.Printf("a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)
}
