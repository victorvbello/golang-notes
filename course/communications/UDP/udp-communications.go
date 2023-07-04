package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
		err = upClientUDP(*addr)
	case TYPE_CNN_SEVER:
		err = upServerUDP(*addr)
	default:
		err = fmt.Errorf("Invalid connection type")
	}

	if err != nil {
		log.Fatal("Error", err)
	}
}

func upClientUDP(addr string) error {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return fmt.Errorf("net.Dial %v", err)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please add message to send")
	for scanner.Scan() {
		fmt.Println("[OUT]>> ", scanner.Text())
		conn.Write(append(scanner.Bytes()))

		serverMsgBff := make([]byte, 1024)
		_, err := conn.Read(serverMsgBff)

		if err != nil {
			if err == io.EOF {
				log.Println("Connection is closed")
				return nil
			}
			return fmt.Errorf("conn.Read %v", err)
		}
		fmt.Println("[IN]>> ", string(serverMsgBff))

		fmt.Println("Please add message to send")
	}
	return nil
}

func upServerUDP(addr string) error {
	pc, err := net.ListenPacket("udp", addr)
	if err != nil {
		return fmt.Errorf("net.ListenPacket %v", err)
	}
	defer pc.Close()
	buffer := make([]byte, 1024)
	fmt.Println("Listening ....")
	for {
		_, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			return fmt.Errorf("pc.ReadFrom %v", err)
		}
		_, err = pc.WriteTo([]byte("Message Received"), addr)
		if err != nil {
			return fmt.Errorf("pc.WriteTo %v", err)
		}
		fmt.Printf("Received %s from address %s \n", string(buffer), addr)
		buffer = buffer[:cap(buffer)]
	}
}
