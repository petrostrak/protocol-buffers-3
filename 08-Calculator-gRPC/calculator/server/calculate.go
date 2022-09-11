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

func (s *Server) CalculatePrimes(in *pb.CalculatorRequest, stream pb.CalculatorService_CalculatePrimesServer) error {
	log.Printf("CalculatePrimes() invoked with %v\n", in)

	var k int32 = 2
	var N int32 = in.A
	for N > 1 {
		if N%k == 0 {
			stream.Send(&pb.CalculatorResponse{
				Result: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}
