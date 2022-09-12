package main

import (
	"log"

	pb "07-grpc-greet-project/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "0.0.0.0:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// greet(c)
	// doGreetManyTimes(c)
	doLongGreet(c)
}
