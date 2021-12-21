package main

import (
	"fmt"
	"runtime"
	"sync"
)

type safeMutexValueInt struct {
	v   int
	mux sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	var count safeMutexValueInt
	var count2 safeMutexValueInt
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			runtime.Gosched()
			count.mux.Lock()
			count2.mux.Lock()
			count.v++
			count2.v += 2
			fmt.Println("Run:", v, "Count:", count.v, "Count 2: ", count2.v)
			count2.mux.Unlock()
			count.mux.Unlock()
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
}
