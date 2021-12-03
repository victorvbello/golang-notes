## Goroutine Scope
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			// var v has the value when goroutine is called
			// var i has the value of current loop flow
			// when this goroutine has executed print the current global value of i
			runtime.Gosched()
			fmt.Println("Run:", v, i)
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
}

```
**Code** https://go.dev/play/p/NBO3Lxe-n90