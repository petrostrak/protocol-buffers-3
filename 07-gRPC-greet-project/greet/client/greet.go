package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"context"
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
