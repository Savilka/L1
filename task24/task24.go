package main

import (
	"L1/task24/point"
	"fmt"
)

func main() {
	a := point.NewPoint(0, 0)
	b := point.NewPoint(1, 1)
	fmt.Println(point.Distance(a, b))
}
