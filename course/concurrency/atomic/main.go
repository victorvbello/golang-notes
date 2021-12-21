package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var count int32

	grs := 100
	wg.Add(grs)
	for i := 0; i < grs; i++ {
		go func(v int) {
			atomic.AddInt32(&count, 2)
			fmt.Println("Run:", v, "Count:", atomic.LoadInt32(&count))
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
}
