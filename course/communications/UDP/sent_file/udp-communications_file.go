package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net"
	"net/http"
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

	fBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return fmt.Errorf("os.ReadFile %v", err)
	}
	fmt.Println("bytes to send ", len(fBytes))
	mimetype := http.DetectContentType(fBytes)
	fmt.Println("mime", mimetype)
	fmt.Println("sending file size:", len(fBytes), " bytes")
	_, err = conn.Write(fBytes)
	if err != nil {
		return fmt.Errorf("conn.Write %v", err)
	}
	return nil
}

func upServerUDP(addr string) error {
	pc, err := net.ListenPacket("udp", addr)
	if err != nil {
		return fmt.Errorf("net.ListenPacket %v", err)
	}
	defer pc.Close()
	// max file size of 4MB
	buffer := make([]byte, 4024)
	fmt.Println("Listening ....")

	n, _, err := pc.ReadFrom(buffer)
	if err != nil {
		return fmt.Errorf("pc.ReadFrom %v", err)
	}
	fmt.Println("bytes received ", n)
	dataReceived := buffer[:n]
	mimetype := http.DetectContentType(dataReceived)
	fmt.Println("mime", mimetype)
	r, err := mime.ExtensionsByType(mimetype)
	if err != nil {
		return fmt.Errorf("mime.ExtensionsByType %v", err)
	}
	fmt.Println(r)
	f, _ := os.Create("received.txt")
	_, err = f.Write(dataReceived)
	if err != nil {
		return fmt.Errorf("file.Write %v", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("file.Close %v", err)
	}
	return nil
}
