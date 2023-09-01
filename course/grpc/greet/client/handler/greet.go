package handler

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/victorvbello/gonotes/course/grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func SentLongGreet(c pb.GreetServiceClient) {
	log.Printf("SentLongGreet request")

	reqs := []*pb.GreetRequest{
		{FirstName: "Victor 1", LastName: "Bello"},
		{FirstName: "Victor 2", LastName: "Bello"},
		{FirstName: "Victor 3", LastName: "Bello"},
		{FirstName: "Victor 4", LastName: "Bello"},
		{FirstName: "Victor 5", LastName: "Bello"},
		{FirstName: "Victor 5", LastName: "Bello"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error on SentLongGreet %v \n", err)
	}

	for _, req := range reqs {
		log.Printf("SentLongGreet req: %v\n", req)

		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error on SentLongGreet stream.CloseAndRecv %v \n", err)
	}

	log.Printf("SentLongGreet: response %s\n", res.Message)
}

func SentMultiGreet(c pb.GreetServiceClient) {
	log.Printf("SentMultiGreet request")

	reqs := []*pb.GreetRequest{
		{FirstName: "Victor 1", LastName: "Bello"},
		{FirstName: "Victor 2", LastName: "Bello"},
		{FirstName: "Victor 3", LastName: "Bello"},
		{FirstName: "Victor 4", LastName: "Bello"},
		{FirstName: "Victor 5", LastName: "Bello"},
		{FirstName: "Victor 5", LastName: "Bello"},
	}

	stream, err := c.MultiGreet(context.Background())

	if err != nil {
		log.Fatalf("Error on SentMultiGreet %v \n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("SentMultiGreet req: %v\n", req)

			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
		log.Printf("stream.CloseSend\n")
	}()

	go func() {
		log.Printf("init stream.Recv\n")
		for {
			res, err := stream.Recv()

			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("Error on SentMultiGreet stream.Recv %v \n", err)
			}

			log.Printf("SentMultiGreet: message %s\n", res.Message)
		}
		close(waitc)
	}()

	<-waitc
}

func SentGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("SentGreetWithDeadline request")

	// waiting time from the client
	cxt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadline(cxt, &pb.GreetRequest{
		FirstName: "Victor",
		LastName:  "Bello",
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("%v\n", e)
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline was exceeded")
				return
			}
			log.Fatalf("Unexpected gRPC err: %v\n", err)
		}
		log.Fatalf("A non Grpc error on SentGreetWithDeadline %v \n", err)
	}

	log.Printf("SentGreetWithDeadline: message %s\n", res.Message)
}
