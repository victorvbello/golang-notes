package main

import (
	"log"
	"time"

	client_handler "github.com/victorvbello/gonotes/course/grpc/greet/client/handler"
	pb "github.com/victorvbello/gonotes/course/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	// by default grpc required SSL
	// add grpc.WithTransportCredentials(insecure.NewCredentials()) for test purpose
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error on grpc.Dial %v on addr %s\n", err, addr)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//client_handler.SentGreet(c)
	//client_handler.SendGreetManyTimes(c)
	//client_handler.SentLongGreet(c)
	//client_handler.SentMultiGreet(c)
	client_handler.SentGreetWithDeadline(c, 5*time.Second)
	client_handler.SentGreetWithDeadline(c, 2*time.Second)
}
