package main

import (
	"fmt"
)

func binSearch(key int, a []int) bool {
	QuickSort(a)

	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		if a[m] == key {
			return true
		}
		if a[m] < key {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

func main() {
	data := []int{7, 5, 1, 1, 99, 2, 2, 3, 4, 5}
	fmt.Println(binSearch(3, data))
}

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
