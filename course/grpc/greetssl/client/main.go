package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/victorvbello/gonotes/course/grpc/greetssl/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const addr = "localhost:50051"

func gRPCLoadClientOptions() ([]grpc.DialOption, error) {
	certFilePath := "course/grpc/greetssl/ssl/server.crt"

	creds, err := credentials.NewClientTLSFromFile(certFilePath, "")
	if err != nil {
		return nil, fmt.Errorf("credentials.NewClientTLSFromFile %v", err)
	}

	return []grpc.DialOption{grpc.WithTransportCredentials(creds)}, nil
}

func main() {

	gRPCClientOpts, err := gRPCLoadClientOptions()

	if err != nil {
		log.Fatalf("Error on gRPCLoadClientOptions %v\n", err)
	}

	conn, err := grpc.Dial(addr, gRPCClientOpts...)

	if err != nil {
		log.Fatalf("Error on grpc.Dial %v on addr %s\n", err, addr)
	}

	defer conn.Close()

	c := pb.NewGreetSSLServiceClient(conn)

	SentGreetSSL(c)

}

func SentGreetSSL(c pb.GreetSSLServiceClient) {
	log.Printf("SentGreetSSL request")

	res, err := c.Greet(context.Background(), &pb.GreetSSLRequest{
		FirstName: "Victor",
		LastName:  "Bello",
	})

	if err != nil {
		log.Fatalf("Error on SentGreetSSL %v \n", err)
	}

	log.Printf("SentGreet: message %s\n", res.Message)
}
