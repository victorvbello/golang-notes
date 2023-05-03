package main

import (
	"testing"
)

func BenchmarkSwitchDefault(b *testing.B) {
	v := 1
	c := 0
	for i := 0; i < b.N; i++ {
		switch v {
		case 2:
			c++
		case 3:
			c++
		case 4:
			c++
		case 5:
			c++
		case 6:
			c++
		case 7:
			c++
		case 8:
			c++
		case 9:
			c++
		case 10:
			c++
		case 11:
			c++
		case 12:
			c++
		default:
			c++
		}
	}
	if c != b.N {
		b.Error("c not equal", c, b.N)
	}
}

func BenchmarkSwitchCase(b *testing.B) {
	v := 1
	c := 0
	for i := 0; i < b.N; i++ {
		switch v {
		case 1:
			c++
		case 2:
			c++
		case 3:
			c++
		case 4:
			c++
		case 5:
			c++
		case 6:
			c++
		case 7:
			c++
		case 8:
			c++
		case 9:
			c++
		case 10:
			c++
		case 11:
			c++
		default:
			c++
		}
	}
	if c != b.N {
		b.Error("c not equal", c, b.N)
	}
}
