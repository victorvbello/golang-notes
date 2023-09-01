package handler

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/victorvbello/gonotes/course/grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet request\n")

	res := ""

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.GreetResponse{
					Message: res,
				})
			}
			log.Fatalf("Error on Receive LongGreet stream.Recv %v \n", err)
		}

		log.Printf("LongGreet receive: %s %s\n", req.FirstName, req.LastName)

		res += fmt.Sprintf("Hi %s %s\n", req.FirstName, req.LastName)
	}
}

func (s *Server) MultiGreet(stream pb.GreetService_MultiGreetServer) error {
	log.Printf("MultiGreet request\n")

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return nil
			}
			log.Fatalf("Error on Receive MultiGreet stream.Recv %v \n", err)
		}

		err = stream.Send(&pb.GreetResponse{
			Message: fmt.Sprintf("Hi %s %s\n", req.FirstName, req.LastName),
		})

		if err != nil {
			log.Fatalf("Error on Sent MultiGreet stream.Send %v \n", err)
		}
	}
}

func (s *Server) GreetWithDeadline(ctx context.Context, r *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline request: %v\n", r)

	// check in context if deadline was exceeded
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client deadline was exceeded")
			// this error is replaced by context deadline error
			return nil, status.Error(codes.Canceled, "The client deadline was exceeded")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Message: "Hi " + r.FirstName + " " + r.LastName,
	}, nil

}
