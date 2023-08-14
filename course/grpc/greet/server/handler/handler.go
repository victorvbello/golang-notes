package handler

import (
	pb "github.com/victorvbello/gonotes/course/grpc/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}
