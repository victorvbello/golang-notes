package main

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"
)

type Person struct {
	First string
	Last  string
}

var p1 Person = Person{"Victor", "Bello"}

func removeTestFile(fileName string) error {
	return os.Remove(fileName)
}

func BenchmarkJSONEncoder(b *testing.B) {
	fileName := "./test_encode.json"
	f, err := os.Create(fileName)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		err = json.NewEncoder(f).Encode(&p1)
		if err != nil {
			b.Error(err)
		}
	}
	f.Close()
	if e := removeTestFile(fileName); e != nil {
		b.Error(e)
	}
}

func BenchmarkWriteFileWithMarshal(b *testing.B) {
	fileName := "./test_write.json"
	f, err := os.Create(fileName)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		writeBuff := bufio.NewWriter(f)
		cb, err := json.Marshal(p1)
		if err != nil {
			b.Error(err)
		}
		writeBuff.WriteString(string(cb))
		if err != nil {
			b.Error(err)
		}
		writeBuff.Flush()
	}
	f.Close()
	if e := removeTestFile(fileName); e != nil {
		b.Error(e)
	}
}
