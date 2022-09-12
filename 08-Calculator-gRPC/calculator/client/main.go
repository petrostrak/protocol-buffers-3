package main

import (
	pb "calculator-grpc-unary-api/calculator/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "0.0.0.0:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to listen on: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// calculate(c)
	// calculatePrimes(c)
	calculateAverage(c)
}
