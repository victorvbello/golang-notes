package main

import "testing"

func factorialLoop(n int) int {
	var f int = 1
	for i := n; i >= 1; i-- {
		f *= i
	}
	return f
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func BenchmarkFactorialLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialLoop(8)
	}
}

func BenchmarkFactorialRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialRecursive(8)
	}
}
