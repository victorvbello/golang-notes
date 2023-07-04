package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	TYPE_CNN_CLIENT = "c"
	TYPE_CNN_SEVER  = "s"
)

func main() {
	var err error

	cnnType := flag.String("type", "", "client(c) / server(s)")
	addr := flag.String("addr", ":3000", "adders (hots:port)")
	flag.Parse()

	switch *cnnType {
	case TYPE_CNN_CLIENT:
		err = upClientTCP(*addr)
	case TYPE_CNN_SEVER:
		err = upServerTCP(*addr)
	default:
		err = fmt.Errorf("Invalid connection type")
	}

	if err != nil {
		log.Fatal("Error", err)
	}
}

func upClientTCP(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Dial %v", err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please add message to send")
	for scanner.Scan() {
		fmt.Println(">> ", scanner.Text())
		conn.Write(append(scanner.Bytes(), '\r'))

		serverMsgBff := make([]byte, 1024)
		_, err := conn.Read(serverMsgBff)

		if err != nil {
			if err == io.EOF {
				log.Println("Connection is closed")
				return nil
			}
			return fmt.Errorf("conn.Read %v", err)
		}
		fmt.Println("<< ", string(serverMsgBff))

		fmt.Println("Please add message to send")
	}
	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner.Err %v", err)
	}
	return nil
}

func upServerTCP(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Listen %v", err)
	}
	defer l.Close()
	log.Println("Server is listening")

	for {
		cnn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("l.Accept %v", err)
		}
		go func(lCnn net.Conn) {
			defer lCnn.Close()
			reader := bufio.NewReader(lCnn)
			writer := bufio.NewWriter(lCnn)
			for {
				lCnn.SetDeadline(time.Now().Add(5 * time.Second))
				rLine, err := reader.ReadString('\r')
				if err != nil {
					if err == io.EOF {
						log.Println("Connection is closed")
						return
					}
					log.Println(fmt.Errorf("reader.ReadString %v", err))
					return
				}
				fmt.Printf("Received %s from address %s \n", rLine[:len(rLine)-1], lCnn.RemoteAddr())
				writer.WriteString("Message received")
				writer.Flush()
			}
		}(cnn)
	}
}
