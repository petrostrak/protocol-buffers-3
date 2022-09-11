package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"context"
	"io"
	"log"
)

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
