package handler

import (
	"context"
	"io"
	"log"

	pb "github.com/victorvbello/gonotes/course/grpc/greet/proto"
)

func SentGreet(c pb.GreetServiceClient) {
	log.Printf("SentGreet request")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Victor",
		LastName:  "Bello",
	})

	if err != nil {
		log.Fatalf("Error on SentGreet %v \n", err)
	}

	log.Printf("SentGreet: message %s\n", res.Message)
}

func SendGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("SendGreetManyTimes request")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Victor",
		LastName:  "Bello",
	})

	if err != nil {
		log.Fatalf("Error on SendGreetManyTimes %v \n", err)
	}

	for {
		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error on SendGreetManyTimes stream.Recv %v \n", err)
		}

		log.Printf("SendGreetManyTimes: message %s\n", res.Message)
	}

}
