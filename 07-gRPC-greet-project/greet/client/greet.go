package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"context"
	"io"
	"log"
)

func greet(c pb.GreetServiceClient) {
	log.Println("greet() was invoked!")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Petros",
	})
	if err != nil {
		log.Printf("Could not greet: %v\n", err)
	}

	log.Printf("Greetin: %s", res.Result)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes() invoked!")

	req := &pb.GreetRequest{
		FirstName: "Petros",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Printf("error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("error while reading stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
