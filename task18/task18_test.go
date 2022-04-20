package main

import "testing"

func Benchmark1(b *testing.B) {
	var c Counter
	for i := 0; i < b.N; i++ {
		c.add()
	}
}

func Benchmark2(b *testing.B) {
	var c CounterMutex
	for i := 0; i < b.N; i++ {
		c.add()
	}
}
