package main

import (
	pb "calculator-grpc-unary-api/calculator/proto"
	"context"
	"log"
)

func (s *Server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculate function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.A + in.B,
	}, nil
}
