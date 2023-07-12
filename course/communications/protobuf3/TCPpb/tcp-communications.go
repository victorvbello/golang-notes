package main

import (
	"bufio"
	"flag"
	"fmt"
	"gonotes/course/communications/protobuf3/hero"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
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
		log.Fatal("Error ", err)
	}
}

func upClientTCP(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Dial %v", err)
	}

	defer conn.Close()

	fmt.Printf("Sending msg\n\n")

	dToSend := &hero.Hero{
		Alias: "Darth vader",
		Name:  "Anakin Skywalker",
		Skills: []*hero.Hero_HeroSkills{
			{Kind: "Lightsaber", Damage: 10.5, Energy: 50},
			{Kind: "Force Healing", Damage: 0.0, Energy: 400},
			{Kind: "Force Lightning", Damage: 50.10, Energy: 300},
		},
	}
	bToSend, err := proto.Marshal(dToSend)

	if err != nil {
		return fmt.Errorf("proto.Marshal %v", err)
	}

	if len(bToSend) == 0 {
		return fmt.Errorf("data to send id empty")
	}

	_, err = conn.Write(bToSend)

	if err != nil {
		return fmt.Errorf("conn.Write %v", err)
	}

	conn.(*net.TCPConn).CloseWrite()

	fmt.Printf("Data send success, total bytes: %d \n\n", len(bToSend))

	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

	serverMsgBff := make([]byte, 1024)
	_, err = conn.Read(serverMsgBff)

	if err != nil {
		if err == io.EOF {
			log.Println("Connection is closed")
			return nil
		}
		return fmt.Errorf("conn.Read %v", err)
	}
	fmt.Printf("<< %s \n\n", string(serverMsgBff))
	return nil
}

func upServerTCP(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Listen %v", err)
	}
	defer l.Close()
	fmt.Printf("Server is listening\n\n")

	for {
		cnn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("l.Accept %v", err)
		}
		fmt.Printf("Cliente connected in %s \n\n", cnn.RemoteAddr())
		go func(lCnn net.Conn) {
			defer lCnn.Close()
			writer := bufio.NewWriter(lCnn)
			for {
				rBytes, err := ioutil.ReadAll(lCnn)
				if err != nil {
					if err == io.EOF {
						log.Println("Connection is closed")
						return
					}
					log.Println(fmt.Errorf("ioutil.ReadAll %v", err))
					return
				}
				if len(rBytes) == 0 {
					continue
				}
				fmt.Printf("Bytes received %d from address %s \n\n", len(rBytes), lCnn.RemoteAddr())

				dReceived := new(hero.Hero)
				err = proto.Unmarshal(rBytes, dReceived)
				if err != nil {
					log.Println(fmt.Errorf("proto.Unmarshal %v", err))
					return
				}

				fmt.Printf("Received data %v from address %s \n\n", dReceived, lCnn.RemoteAddr())
				writer.WriteString("Message received")
				writer.Flush()
			}
		}(cnn)
	}
}
