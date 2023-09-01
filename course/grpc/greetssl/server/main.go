package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/victorvbello/gonotes/course/grpc/greetssl/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const addr = "127.0.0.1:50051"

type ServerSSL struct {
	pb.GreetSSLServiceServer
}

func (s *ServerSSL) Greet(ctx context.Context, r *pb.GreetSSLRequest) (*pb.GreetSSLResponse, error) {
	log.Printf("Greet request: %v\n", r)
	return &pb.GreetSSLResponse{
		Message: fmt.Sprintf("Hi %s %s", r.FirstName, r.LastName),
	}, nil
}

func gRPCLoadServerOptions() ([]grpc.ServerOption, error) {
	certFilePath := "course/grpc/greetssl/ssl/server.crt"
	keyFilePath := "course/grpc/greetssl/ssl/server.pem"

	creds, err := credentials.NewServerTLSFromFile(certFilePath, keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("credentials.NewServerTLSFromFile %v", err)
	}

	return []grpc.ServerOption{grpc.Creds(creds)}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error on net.Listen %v\n", err)
	}

	log.Printf("Greet Listening on %s\n", addr)

	gRPCServerOpts, err := gRPCLoadServerOptions()

	if err != nil {
		log.Fatalf("Error on gRPCLoadServerOptions %v\n", err)
	}

	s := grpc.NewServer(gRPCServerOpts...)

	pb.RegisterGreetSSLServiceServer(s, &ServerSSL{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error on s.Server %v\n", err)
	}

}
