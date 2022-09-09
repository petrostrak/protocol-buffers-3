package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"context"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Panicf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
