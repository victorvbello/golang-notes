package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gonotes/course/concurrency/pipelines/chat"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	clientName := fmt.Sprintf("Unknow:%d", rand.Intn(100))
	fmt.Println("[CHAT][INIT]")
	fmt.Println("Please add your name")
	fmt.Scanln(&clientName)

	log.Printf("[CHAT][INIT][CONNECTING] user:%s \n", clientName)
	conn, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		log.Fatalf("[CHAT][ERROR][net.Dial] connecting to chat, %v", err.Error())
	}
	log.Printf("[CHAT][CONNECTED]")
	defer conn.Close()

	chatConfig := chat.ChatClientConfig{
		Name: clientName,
	}

	config, err := json.Marshal(chatConfig)
	if err != nil {
		log.Fatalf("[CHAT][ERROR][json.Marshal] chat config, %v", err.Error())
	}

	// Set chat config to server
	_, err = fmt.Fprint(conn, chat.CHAT_CONFIG_FLAG, string(config), chat.CHAT_CONFIG_FLAG+"\n")
	if err != nil {
		log.Fatalf("[CHAT][ERROR][Fprint] send config, %v", err.Error())
	}

	// Flow input msg
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			msg := scanner.Text()
			if strings.Contains(msg, fmt.Sprintf(chat.USER_MSG_LAYOUT, clientName)) {
				continue
			}
			fmt.Println(msg)
		}
	}()

	// Flow output msg
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprint(conn, msg+"\n")
		if err != nil {
			log.Printf("[CHAT][ERROR][Fprint] send config, %v", err.Error())
		}
	}
}
