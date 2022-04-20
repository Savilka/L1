package main

import "fmt"

func main() {
	data := [5]interface{}{1, 3.14, false, make(chan interface{}), "string"}
	for _, v := range data {
		switch v.(type) {
		case int:
			fmt.Printf("%d is integer\n", v)
		case string:
			fmt.Printf("%s is string\n", v)
		case bool:
			fmt.Printf("%t is bool\n", v)
		case chan interface{}:
			fmt.Printf("%v is channel(pointer)\n", v)
		default:
			fmt.Println("unsupported type")
		}
	}
}
