package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s/%s", "1", "2")
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "1" + "/" + "2"
	}
}

func BenchmarkJoin(b *testing.B) {
	xs := []string{"1", "/", "2"}
	for i := 0; i < b.N; i++ {
		_ = strings.Join(xs, "")
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bb bytes.Buffer
		bb.WriteString("1")
		bb.WriteString("/")
		bb.WriteString("2")
		_ = bb.String()
	}
}
