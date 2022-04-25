```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() { //sent to chan c1
		c1 <- "hi from c1"
	}()

	go func() { //sent to chan c2
		time.Sleep(1 * time.Second)
		c2 <- "hi from c2"
	}()

	for {
		select {
		case m := <-c1:
			fmt.Println("Received from c1,", m)
		case m := <-c2:
			fmt.Println("Received from c2,", m)
		case <-time.After(2 * time.Second):
			fmt.Println("flow timeout")
			return
		}
	}

}
```
**Code** https://go.dev/play/p/19kA2CLKm32