package main

import (
	pb "calculator-grpc-unary-api/calculator/proto"
	"context"
	"io"
	"log"
	"time"
)

func calculate(c pb.CalculatorServiceClient) {
	log.Println("calculate() was invoked!")
	resp, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		A: 3,
		B: 9,
	})
	if err != nil {
		log.Printf("could not calculate: %v\n", err)
	}

	log.Printf("The sum is: %d", resp.Result)
}

func calculatePrimes(c pb.CalculatorServiceClient) {
	log.Println("calculatePrimes() invoked!")

	req := &pb.CalculatorRequest{
		A: 120,
	}

	stream, err := c.CalculatePrimes(context.Background(), req)
	if err != nil {
		log.Printf("error while calling CalculatePrimes: %v", err)
	}

	for {
		prime, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("error while reading stream: %v\n", err)
		}

		log.Printf("CalculatePrimes: %d\n", prime.Result)
	}
}

func calculateAverage(c pb.CalculatorServiceClient) {
	log.Println("calculatePrimes() invoked!")

	requests := []*pb.CalculatorRequest{
		{A: 3},
		{A: 9},
		{A: 12},
		{A: 31},
		{A: 4},
	}

	stream, err := c.CalculateAverage(context.Background())
	if err != nil {
		log.Printf("error while calling CalculateAverage: %v\n", err)
	}

	for _, req := range requests {
		log.Printf("sending number: %d\n", req.A)

		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("error while receiving response from CalculateAverage: %v\n", err)
	}

	log.Printf("CalculateAverage: %d\n", res.Result)
}
