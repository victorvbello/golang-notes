package chat

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const CHAT_CONFIG_FLAG = "###CHAT_CONFIG###"
const USER_MSG_LAYOUT = "<%s>"
const MSG_LAYOUT = USER_MSG_LAYOUT + "%s"

type ChatClientConfig struct {
	Name string `json:"name"`
}

func Run(connection string) error {
	l, err := net.Listen("tcp", connection)
	if err != nil {
		log.Println("[CHAT][ERROR][net.Listen] connecting to chan", connection, err.Error())
		return err
	}
	r := CreateRoom("chatRoom")

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		l.Close()
		log.Println("[CHAT] tcp connection closed")
		close(r.quit)
		if r.TotalClients() > 0 {
			<-r.msgChan
		}
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[CHAT][ERROR][l.Accept] accepting connection from chat client", err.Error())
			break
		}
		go handleConnection(r, conn)
	}

	return err
}

func handleConnection(r *room, conn net.Conn) {
	log.Println("[CHAT][handleConnection] received request from client", conn.RemoteAddr())
	r.AddClient(conn)
}
