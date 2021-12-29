package main

import (
	"fmt"

	"./mystring"
)

func main() {
	fmt.Println(mystring.MyStringJoin("hi", "maria"))
	fmt.Println(mystring.MyStringJoin("hi", "luis", "how", "are", "you"))
}
