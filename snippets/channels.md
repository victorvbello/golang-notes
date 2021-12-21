## Channels

This code show how to use channel control flow

```go
package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	fmt.Println("send 1")
	time.Sleep(1 * time.Second)

	c <- 42
	fmt.Println("send 2")
}

func receive(c <-chan int) {
	fmt.Println("receive 1")
	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	fmt.Println("receive 2")
}

func main() {
	c := make(chan int)

	go send(c) // func send change the channel type to (chan<- int), this only permit push value to the channel

	receive(c) // func receive change the channel type to (<-chan int), this only permit pull value from the channel

	fmt.Println("end")
}

```
**Code:** https://go.dev/play/p/jXM_D4IpCam.go