package main

import "fmt"

// go test -bench . -benchmem
// Benchmark1-8     2965014               394.7 ns/op            80 B/op          8 allocs/op
// Benchmark2-8     9576471               117.5 ns/op             0 B/op          0 allocs/op
// Из бенчмарков видно, что первый способ уступает как по памяти, так и по времени выполнения

func Reverse1(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func Reverse2(s string) string {
	runeS := []rune(s)
	for i := 0; i < len(runeS)/2; i++ {
		runeS[i], runeS[len(runeS)-i-1] = runeS[len(runeS)-i-1], runeS[i]
	}

	return string(runeS)
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(Reverse1(str))
	fmt.Println(Reverse2(str))
}
