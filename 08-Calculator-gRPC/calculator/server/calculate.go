package main

import (
	pb "calculator-grpc-unary-api/calculator/proto"
	"context"
	"io"
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

func (s *Server) CalculateAverage(stream pb.CalculatorService_CalculateAverageServer) error {
	log.Println("CalculateAverage() invoked")

	var sum int32 = 0
	var counter int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CalculatorResponse{
				Result: sum / counter,
			})
		} else if err != nil {
			log.Printf("error while reading client stream: %v\n", err)
		}

		sum += req.A
		counter++
		log.Printf("Received: %d\n", req.A)
		log.Printf("Sum: %d", sum)
		log.Printf("Counter: %d", counter)
	}
}

func (s *Server) CalculateMax(stream pb.CalculatorService_CalculateMaxServer) error {
	log.Println("CalculateMax() invoked")

	var max int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Max number was: %d\n", max)
			return nil
		} else if err != nil {
			log.Printf("error while reading client stream: %v\n", err)
		}

		if req.A > max {
			max = req.A
		}

		err = stream.Send(&pb.CalculatorResponse{
			Result: max,
		})
		if err != nil {
			log.Printf("error while sending data to client: %v\n", err)
		}
	}
}
