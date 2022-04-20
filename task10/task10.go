package main

import "fmt"

func main() {
	data := [15]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 2, 6, 0, -4, -9, 10, -10}
	group := make(map[int][]float64)

	for _, t := range data {
		if t < 0 {
			group[(int(t/10)-1)*10] = append(group[(int(t/10)-1)*10], t)
		} else {
			group[int(t/10)*10] = append(group[int(t/10)*10], t)
		}
	}
	fmt.Println(group)
}
