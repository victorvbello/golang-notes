package main

import (
	"log"
	"os"
	"time"
)

func main() {
	filePath := "../test.txt"
	f, err := os.Stat(filePath)
	if err != nil {
		log.Fatal("Error os.Stat", err)
	}
	for {
		time.Sleep(1 * time.Second)
		fc, err := os.Stat(filePath)
		if err != nil {
			log.Fatal("Error os.Stat", err)
		}
		if f.ModTime() != fc.ModTime() {
			log.Printf("file %s change %s", filePath, fc.ModTime())
			f, err = os.Stat(filePath)
			if err != nil {
				log.Fatal("Error os.Stat", err)
			}
		}
	}
}
