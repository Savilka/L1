package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	var s string
	fmt.Scan(&s)

	bigA := big.NewFloat(a)
	bigB := big.NewFloat(b)

	switch s {
	case "+":
		sum := new(big.Float)
		sum = sum.Add(bigA, bigB)
		fmt.Println(sum.String())
	case "-":
		sub := new(big.Float)
		sub = sub.Sub(bigA, bigB)
		fmt.Println(sub.String())
	case "*":
		mul := new(big.Float)
		mul = mul.Mul(bigA, bigB)
		fmt.Println(mul.String())
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
		} else {
			div := new(big.Float)
			div = div.Quo(bigA, bigB)
			fmt.Println(div.String())
		}
	default:
		fmt.Println("bad operation")
	}

}
