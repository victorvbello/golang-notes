package itoa_vs_sprintf

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	testValueInt1 int = iota + 5647
	testValueInt2 int = iota + testValueInt1
	testValueInt3 int = iota + testValueInt1
	testValueInt4 int = iota + testValueInt1
	testValueInt5 int = iota + testValueInt1
)

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(testValueInt1)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", testValueInt1)
	}
}

func BenchmarkItoaMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(testValueInt1)
		_ = strconv.Itoa(testValueInt2)
		_ = strconv.Itoa(testValueInt3)
		_ = strconv.Itoa(testValueInt4)
		_ = strconv.Itoa(testValueInt5)
	}
}

func BenchmarkSprintfMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d, %d, %d, %d, %d",
			testValueInt1,
			testValueInt2,
			testValueInt3,
			testValueInt4,
			testValueInt5)
	}
}
