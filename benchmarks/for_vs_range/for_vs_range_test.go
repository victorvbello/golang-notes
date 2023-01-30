package main

import (
	"testing"
)

var x1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func BenchmarkFor(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		for xx := 0; xx < len(x1); xx++ {
			x += x1[xx]
		}
	}
}

func BenchmarkRange(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		for _, xx := range x1 {
			x += xx
		}
	}
}
