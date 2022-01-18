package main

import (
	"encoding/json"
	"testing"
)

type Person struct {
	First string
	Last  string
}

type LocalWriter struct{}

func (lw *LocalWriter) Write([]byte) (int, error) {
	return 0, nil
}

var p1 Person = Person{"Victor", "Bello"}

var Lw *LocalWriter = &LocalWriter{}

func BenchmarkEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.NewEncoder(Lw).Encode(p1)
	}
}

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b, _ := json.Marshal(p1)
		Lw.Write(b)
	}
}
