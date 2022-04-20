package main

import "fmt"

// Алгоритм пересечения двух множеств

func main() {
	set1 := []int{1, 5, 7, 2, 3}
	set2 := []int{5, 9, 1, 343, 7}

	intersectionMap := make(map[int]int, len(set1))
	var intersection []int
	for _, val := range set1 {
		intersectionMap[val] = val
	}

	for _, val := range set2 {
		_, ok := intersectionMap[val]
		if ok {
			intersection = append(intersection, val)
		}
	}

	fmt.Println(intersection)
}
