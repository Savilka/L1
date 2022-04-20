package main

import "testing"

func Benchmark1(b *testing.B) {
	str := "главрыба"
	for i := 0; i < b.N; i++ {
		Reverse1(str)
	}
}

func Benchmark2(b *testing.B) {
	str := "главрыба"
	for i := 0; i < b.N; i++ {
		Reverse2(str)
	}
}
