package main

import (
	"context"
	"io"
	"log"

	pb "github.com/victorvbello/gonotes/course/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "127.0.0.1:50052"

func RequestSum(c pb.CalculatorServiceClient) {
	log.Printf("RequestSum request")

	sumRequest := &pb.SumRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Sum(context.Background(), sumRequest)

	if err != nil {
		log.Fatalf("Error on RequestSum %v \n", err)
	}

	log.Printf("RequestSum: %d + %d = %d\n", sumRequest.A, sumRequest.B, res.Result)
}

func RequestPrimes(c pb.CalculatorServiceClient) {
	log.Printf("RequestPrimes request")

	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Error on Primes %v \n", err)
	}

	for {
		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error on Primes stream.Recv %v \n", err)
		}

		log.Printf("Primes: result %d\n", res.Result)
	}
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error on grpc.Dial %v on addr %s\n", err, addr)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	RequestSum(c)
	RequestPrimes(c)
}
