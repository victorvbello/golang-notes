## Array vs Slice
``` go
package main

import (
	"fmt"
)

// Diff into Arrays and Slices

func exampleArray() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered of", r)
		}
	}()
	a := [5]int{1, 2, 3, 4, 5} //Array
	fmt.Printf("Array a: %v len: %d cap: %d type: %T\n", a, len(a), cap(a), a)
	forceAddIntoArray(a, 6) // this failt because array only have 5 positions
}

func forceAddIntoArray(a [5]int, i int) {
	a[i] = i
}

func main() {
	exampleArray()
	s := []int{1, 2, 3, 4, 5} // Slice literal
	fmt.Printf("Slice s: %v len: %d cap: %d type: %T\n", s, len(s), cap(s), s)
	s1 := make([]int, 5, 5) // Slice make
	fmt.Printf("Slice s1: %v len: %d cap: %d type: %T\n", s1, len(s1), cap(s1), s1)
	for i, _ := range s1 { // fill slice
		s1[i] = i + 1
	}
	fmt.Printf("Slice s1: %v len: %d cap: %d type: %T\n", s1, len(s1), cap(s1), s1)
	s1 = append(s1, 1)
	fmt.Printf("Slice s1: %v len: %d cap: %d type: %T\n", s1, len(s1), cap(s1), s1) // if you use make consider this
}
```
**Code:** https://play.golang.org/p/9dtGeK94TEg