package main

import (
	"gonotes/course/http/client"
	"gonotes/course/http/server"
)

func main() {
	/* 	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Test http end")
		os.Exit(0)
	}() */

	go func() {
		client.Start()
	}()

	server.Start()
}
