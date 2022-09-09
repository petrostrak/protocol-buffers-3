package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
