## Interface implementation

We have two types that implement the shape interface, because but have de method `area() float32`

```go
// You can edit this code!
// Click here and start typing.
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

func info(sh shape) {
	fmt.Printf("the area is %.2f of type %T\n", sh.area(), sh)
}

type squared struct {
	length int
	width  int
}

func (s squared) area() float32 {
	return float32(s.length * s.width)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius
}

func main() {
	s := squared{
		length: 2,
		width:  3,
	}
	c := circle{
		radius: 4,
	}

	info(s)
	info(c)

}
```