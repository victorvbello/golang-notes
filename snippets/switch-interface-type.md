```go
package main

import "fmt"

type Human interface {
	Speak(s string)
}

type Jon struct{}

func (j Jon) Speak(s string) {
	fmt.Println("Jon say ", s)
}

type Susan struct{}

func (s Susan) Speak(s string) {
	fmt.Println("Susan say ", s)
}

func checkType(h Human) {
	switch c := h.(type) {
	case *Jon:
		fmt.Println("--- type is Jon")
		c.Speak("hi")
	default:
		fmt.Printf("--- type %T without case\n", c)
	}
}

func main() {
	var h Human
	h = &Jon{}
	checkType(h)
	h = &Susan{}
	checkType(h)

}
```
**Code** https://go.dev/play/p/VU9NWBgtG6x