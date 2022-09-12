package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"context"
	"io"
	"log"
	"time"
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

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet() invoked!")

	reqs := []*pb.GreetRequest{
		{FirstName: "Petros"},
		{FirstName: "Eirini"},
		{FirstName: "Maggie"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
