package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/victorvbello/gonotes/course/grpc/calculator/proto"

	"google.golang.org/grpc"
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

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error on net.Listen %v\n", err)
	}

	log.Printf("Calculator Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error on s.Server %v\n", err)
	}
}
