``` go
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := s1[3:9]
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	s2[0] = 999
	fmt.Println("s2 with the 0 position changed", s2[0])
	fmt.Println("s1 with the 3 position changed", s1[3])

}
```
**Code** https://go.dev/play/p/psJgXKsp4xq