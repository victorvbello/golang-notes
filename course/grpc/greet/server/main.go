package main

import (
	pb "github.com/victorvbello/gonotes/course/grpc/greet/proto"
	server_handler "github.com/victorvbello/gonotes/course/grpc/greet/server/handler"
	"log"
	"net"

	"google.golang.org/grpc"
)

const addr = "127.0.0.1:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error on net.Listen %v\n", err)
	}

	log.Printf("Greet Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &server_handler.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error on s.Server %v\n", err)
	}

}
