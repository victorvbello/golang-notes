package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	pb "github.com/victorvbello/gonotes/course/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

func RequestAvg(c pb.CalculatorServiceClient) {
	log.Printf("RequestAvg request")

	inputsString := []string{}

	inputs := []int64{1, 2, 3, 4}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error on RequestAvg %v \n", err)
	}

	for _, n := range inputs {

		inputsString = append(inputsString, fmt.Sprintf("%d", n))

		stream.Send(&pb.AvgRequest{
			Number: n,
		})

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error on RequestAvg stream.CloseAndRecv %v \n", err)
	}

	log.Printf("RequestAvg: response (%s)/%d = %.2f\n", strings.Join(inputsString, "+"), len(inputsString), res.Result)

}

func RequestMax(c pb.CalculatorServiceClient) {
	log.Printf("RequestMax request")

	var inputs, outputs []int32

	inputs = []int32{1, 5, 3, 6, 2, 20}

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error on RequestMax %v \n", err)
	}

	quit := make(chan struct{})

	go func() {
		for _, n := range inputs {
			log.Printf("SentMax req: %d\n", n)

			stream.Send(&pb.MaxRequest{
				Number: n,
			})

			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
		log.Printf("SentMax stream.CloseSend\n")
	}()

	go func() {
		log.Printf("SentMax init stream.Recv\n")
		for {
			res, err := stream.Recv()

			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("Error on SentMax stream.Recv %v \n", err)
			}

			log.Printf("SentMax: result %d\n", res.Result)
			outputs = append(outputs, res.Result)
		}
		close(quit)
	}()

	<-quit

	log.Printf("SentMax: final result %v = %v\n", inputs, outputs)
}

func RequestSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Printf("RequestSqrt request")

	sqrtRequest := &pb.SqrtRequest{
		Number: n,
	}

	res, err := c.Sqrt(context.Background(), sqrtRequest)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("gRPC error, code: %v message: %v\n", e.Code(), e.Message())
			if e.Code() == codes.InvalidArgument {
				log.Printf("Sent a invalid argument")
				return
			}
		} else {
			log.Fatalf("not gRPC error %v\n", err)
		}

	}

	log.Printf("RequestSqrt: Sqrt(%d) = %.02f\n", sqrtRequest.Number, res.Result)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error on grpc.Dial %v on addr %s\n", err, addr)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//RequestSum(c)
	//RequestPrimes(c)
	//RequestAvg(c)
	//RequestMax(c)
	RequestSqrt(c, 10)
	RequestSqrt(c, -10)
}
