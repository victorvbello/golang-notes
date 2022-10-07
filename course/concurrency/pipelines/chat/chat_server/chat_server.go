package main

import "gonotes/course/concurrency/pipelines/chat"

func main() {
	chat.Run("127.0.0.1:4000")
}
