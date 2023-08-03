package server

import "golang.org/x/net/websocket"

var CHAT_MSG = make(chan string)

var CHAT_CLIENTS = make(map[int]*websocket.Conn)
