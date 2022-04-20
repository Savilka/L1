package main

import "fmt"

func deleteElement(a []int, i int) []int {
	if i < len(a) && i >= 0 {
		return append(a[0:i], a[i+1:]...)
	}
	fmt.Println("Bad index")
	return nil
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data = deleteElement(data, 3)
	fmt.Println(data)
}
