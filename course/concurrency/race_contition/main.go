package main

import (
	"fmt"
	"sync"
)

// go run -race .course/race_contition/main.go
// Found 2 data race(s)

func main() {
	var wg sync.WaitGroup
	var count int
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			count++
			fmt.Println("Run:", v, i, "Count:", count)
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
}
