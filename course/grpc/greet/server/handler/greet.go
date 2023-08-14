package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/victorvbello/gonotes/course/grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, r *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet request: %v\n", r)
	return &pb.GreetResponse{
		Message: fmt.Sprintf("Hi %s %s", r.FirstName, r.LastName),
	}, nil
}

func (s *Server) GreetManyTimes(r *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes request: %v\n", r)

	for i := 1; i <= 10; i++ {
		res := fmt.Sprintf("- %d Hi %s %s", i, r.FirstName, r.LastName)
		stream.Send(&pb.GreetResponse{
			Message: res,
		})
		time.Sleep(1 * time.Second)
	}

	return nil
}
