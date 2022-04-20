package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	data := [5]int{2, 4, 6, 8, 10}

	for i := 0; i < 5; i++ {
		// Вычисляем квадрат числа в отдельной горутине
		go func(a int) {
			fmt.Fprintln(os.Stdout, a*a)
		}(data[i])
	}

	// Sleep чтобы main не завершился раньше чем все горутины посчитают квадраты
	time.Sleep(time.Second)
}
