package mystring

import (
	"fmt"
	"testing"
)

func TestMyStringJoin(t *testing.T) {
	ex := "hi,maria"
	r := MyStringJoin("hi", "maria")
	if r != ex {
		t.Errorf("Expected %s Got %s", ex, r)
	}
}

func TestTableMyStringJoin(t *testing.T) {
	tcs := []struct {
		name     string
		data     []string
		expected string
	}{
		{"Join numbers", []string{"1", "2", "3"}, "1,2,3"},
		{"Join numbers", []string{"A", "B", "C"}, "A,B,C"},
		{"Join numbers", []string{"maria", "luis", "pedro"}, "maria,luis,pedro"},
	}
	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			r := MyStringJoin(c.data...)
			if r != c.expected {
				t.Errorf("\nCase: %s\n\tExpected: \t%s\n\tGot: \t\t%s", t.Name(), c.expected, r)
			}
		})
	}
}

func ExampleMyStringJoin() {
	fmt.Println(MyStringJoin("hi", "maria"))
	fmt.Println(MyStringJoin("hi", "pedro"))
	// Output:
	// hi,maria
	// hi,pedro
}

func BenchmarkMyStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyStringJoin("hi", "maria")
	}
}
