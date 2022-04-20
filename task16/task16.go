package main

import "fmt"

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	i, j := 0, len(a)-1
	p := a[len(a)>>1]
	for i <= j {
		for a[i] < p {
			i++
		}
		for a[j] > p {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	if j > 0 {
		QuickSort(a[:j+1])
	}
	if len(a) > i {
		QuickSort(a[i:])
	}
}

func main() {
	data := []int{4, 6, 1, 8, 2, 65, 1, -2, 0}
	QuickSort(data)
	fmt.Println(data)
}
