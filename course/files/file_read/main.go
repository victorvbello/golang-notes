package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../test.txt")
	if err != nil {
		log.Fatal("Error os.Open", err)
	}
	scanner := bufio.NewScanner(f)
	line := 0
	for scanner.Scan() {
		line++
		fmt.Printf("- line:%d\t content:%s\n", line, scanner.Text())
	}
}
