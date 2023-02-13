package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// clear current content
	if err := os.Truncate("./test.txt", 0); err != nil {
		log.Fatal("Error os.Truncate", err)
	}

	f, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error os.Open", err)
	}
	writeBuff := bufio.NewWriter(f)

	lines := []string{"hi", "", "how", "are", "you"}
	for i, l := range lines {
		_, err := writeBuff.WriteString(fmt.Sprintf("line:%d\t%s\n", i, l))
		if err != nil {
			log.Fatal("Error writeBuff.WriteString", err)
		}
	}
	err = writeBuff.Flush()
	if err != nil {
		log.Fatal("Error writeBuff.Flush", err)
	}
}
