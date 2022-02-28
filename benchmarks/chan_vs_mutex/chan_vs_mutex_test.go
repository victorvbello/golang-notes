package main

import (
	"fmt"
	"sync"
	"testing"
)

type safeMutexValueString struct {
	v   []string
	mux sync.Mutex
}

func BenchmarkChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var count int
		var wg sync.WaitGroup
		tc := 100
		c := make(chan string)
		for x := 0; x < tc; x++ {
			wg.Add(1)
			go func(index int) {
				c <- fmt.Sprintf("Chan-%d", index)
				wg.Done()
			}(x)
		}

		go func() {
			wg.Wait()
			close(c)
		}()

		for s := range c {
			count++
			_ = fmt.Sprint(s)
		}
		// check process all cases
		if count != tc {
			b.Error("flow not test all cases")
		}
	}
}

func BenchmarkMutexSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var count int
		var c safeMutexValueString
		var wg sync.WaitGroup
		tc := 100
		for x := 0; x < tc; x++ {
			wg.Add(1)
			go func(index int) {
				c.mux.Lock()
				c.v = append(c.v, fmt.Sprintf("Mutex-%d", index))
				c.mux.Unlock()
				wg.Done()
			}(x)
		}

		wg.Wait()

		c.mux.Lock()
		for _, s := range c.v {
			count++
			_ = fmt.Sprint(s)
		}
		c.mux.Unlock()
		// check process all cases
		if count != tc {
			b.Error("flow not test all cases")
		}
	}
}

func BenchmarkMutexArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var count int
		var c safeMutexValueString
		var wg sync.WaitGroup
		tc := 100
		c.v = make([]string, tc)
		for x := 0; x < tc; x++ {
			wg.Add(1)
			go func(index int) {
				c.mux.Lock()
				c.v[index] = fmt.Sprintf("Mutex-%d", index)
				c.mux.Unlock()
				wg.Done()
			}(x)
		}

		wg.Wait()

		c.mux.Lock()
		for _, s := range c.v {
			count++
			_ = fmt.Sprint(s)
		}
		c.mux.Unlock()
		// check process all cases
		if count != tc {
			b.Error("flow not test all cases")
		}
	}
}
