package main

import (
	_ "net/http/pprof"
)

// go test -bench . -benchmem -cpuprofile cpu.out -memprofile mem.out -memprofilerate=1
// go tool pprof  task15.test.exe mem.out
//
// Когда someFunc() завершила работу, сборщик мусора не отчистил память большой строки, так как justString1 продолжала
// ссылаться на v. Проблему можно исправить, если скопировать нужные нам данные в новый слайс. Тогда сборщик мусора отчистит
// всю ненужную память.
// Результаты профилировщика
// 1kB     44:   _ = someFunc()
// 64B     48:   _ = mySomeFunc()
// Разница очевидна

var justString1 string
var justString2 string

func someFunc() string {
	v := createHugeString(1 << 10)
	justString1 = v[:50]
	return justString1
}

func mySomeFunc() string {
	v := createHugeString(1 << 10)
	buf := make([]byte, 50)
	copy(buf, v[:50])
	justString2 = string(buf)
	return justString2
}

func createHugeString(i int) string {
	hugeSliceOfBytes := make([]byte, i)
	for j := 0; j < i; j++ {
		hugeSliceOfBytes[j] = 44
	}
	return string(hugeSliceOfBytes)
}

func test1() {
	_ = someFunc()
}

func test2() {
	_ = mySomeFunc()
}

func main() {

}
