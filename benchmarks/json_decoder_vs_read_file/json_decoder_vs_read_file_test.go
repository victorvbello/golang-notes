package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Person struct {
	First string
	Last  string
}

const (
	fileName = "./test.json"
)

func createTestFile() (*os.File, error) {
	p1 := Person{"Victor", "Bello"}
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	writeBuff := bufio.NewWriter(f)
	cb, err := json.Marshal(p1)
	if err != nil {
		return nil, err
	}
	_, err = writeBuff.WriteString(string(cb))
	if err != nil {
		return nil, err
	}
	err = writeBuff.Flush()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func openTestFile() (*os.File, error) {
	return os.Open(fileName)
}

func removeTestFile() error {
	return os.Remove(fileName)
}

func BenchmarkJSONDecoder(b *testing.B) {
	_, err := createTestFile()
	if err != nil {
		b.Error(fmt.Errorf("createTestFile %v", err))
	}
	for i := 0; i < b.N; i++ {

		var p1 Person
		f, err := openTestFile()
		if err != nil {
			b.Error(fmt.Errorf("openTestFile %v", err))
		}
		if err = json.NewDecoder(f).Decode(&p1); err != nil {
			b.Error(fmt.Errorf("NewDecoder.Decode %v", err))
		}

		if p1.First != "Victor" {
			b.Errorf("Fist is not equal, %s", p1.First)
		}
		f.Close()
	}
	if e := removeTestFile(); e != nil {
		b.Error(e)
	}
}

func BenchmarkReadFileWithUnmarshal(b *testing.B) {
	_, err := createTestFile()
	if err != nil {
		b.Error(fmt.Errorf("createTestFile %v", err))
	}
	for i := 0; i < b.N; i++ {

		var p1 Person
		var allContent []byte
		f, err := openTestFile()
		if err != nil {
			b.Error(fmt.Errorf("openTestFile %v", err))
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			allContent = append(allContent, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			b.Error(fmt.Errorf("scanner %v", err))
		}
		if err := json.Unmarshal(allContent, &p1); err != nil {
			b.Error(fmt.Errorf("json.Unmarshal %v", err))
		}

		if p1.First != "Victor" {
			b.Errorf("Fist is not equal, %s", p1.First)
		}
		f.Close()
	}
	if e := removeTestFile(); e != nil {
		b.Error(e)
	}
}
