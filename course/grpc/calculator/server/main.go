package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	pb "github.com/victorvbello/gonotes/course/grpc/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const addr = "127.0.0.1:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func (s *Server) Sum(ctx context.Context, r *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum request: %v\n", r)
	return &pb.SumResponse{
		Result: r.A + r.B,
	}, nil
}

func (s *Server) Primes(r *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes request: %v\n", r)

	k := int64(2)
	N := r.Number

	for N > 1 {
		if N%k != 0 {
			k++
			continue
		}

		N /= k

		stream.Send(&pb.PrimesResponse{
			Result: k,
		})

		time.Sleep(1 * time.Second)
	}

	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg request\n")

	var sum, count int64
	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error on Receive Avg stream.Recv %v \n", err)
		}
		sum += req.Number
		count++
	}

	return stream.SendAndClose(&pb.AvgResponse{
		Result: float64(sum) / float64(count),
	})
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max request\n")
	var max int32

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			log.Fatalf("Error on Receive Max stream.Recv %v \n", err)
		}

		if n := req.Number; n > max {
			max = n
			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error on Sent Max stream.Send %v \n", err)
			}
		}

	}
}

func (s *Server) Sqrt(ctx context.Context, r *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt request: %v\n", r.Number)

	number := r.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Request number %d is invalid, number should be > 0", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error on net.Listen %v\n", err)
	}

	log.Printf("Calculator Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error on s.Server %v\n", err)
	}
}
