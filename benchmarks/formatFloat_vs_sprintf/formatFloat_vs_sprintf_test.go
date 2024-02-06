package formatFloat_vs_sprintf

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	testValueFloat64_1 float64 = iota + 3.1415926535
	testValueFloat64_2 float64 = iota + testValueFloat64_1
	testValueFloat64_3 float64 = iota + testValueFloat64_1
	testValueFloat64_4 float64 = iota + testValueFloat64_1
	testValueFloat64_5 float64 = iota + testValueFloat64_1
)

func BenchmarkFormatFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatFloat(testValueFloat64_1, 'g', -1, 64)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", testValueFloat64_1)
	}
}

func BenchmarkFormatFloatMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatFloat(testValueFloat64_1, 'g', -1, 64)
		_ = strconv.FormatFloat(testValueFloat64_2, 'g', -1, 64)
		_ = strconv.FormatFloat(testValueFloat64_3, 'g', -1, 64)
		_ = strconv.FormatFloat(testValueFloat64_4, 'g', -1, 64)
		_ = strconv.FormatFloat(testValueFloat64_5, 'g', -1, 64)
	}
}

func BenchmarkSprintfMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v, %v, %v, %v, %v",
			testValueFloat64_1,
			testValueFloat64_2,
			testValueFloat64_3,
			testValueFloat64_4,
			testValueFloat64_5)
	}
}
