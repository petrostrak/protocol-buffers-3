package main

import (
	pb "calculator-grpc-unary-api/calculator/proto"
	"context"
	"log"
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
