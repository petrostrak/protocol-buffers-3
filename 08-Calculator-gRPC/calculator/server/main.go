package main

import (
	pb "calculator-grpc-unary-api/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	err = s.Serve(l)
	if err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
