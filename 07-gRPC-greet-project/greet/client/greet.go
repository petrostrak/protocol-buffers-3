package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone() invoked!")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Printf("error while creating steam: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Petros"},
		{FirstName: "Eirini"},
		{FirstName: "Maggie"},
	}

	waitChan := make(chan struct{})

	// This goroutine sends the requests to the server
	go func() {
		for _, req := range reqs {
			log.Printf("send request: %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Printf("error sending request: %v\n", err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	// This goroutine waits for the response from the server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("error while receiving: %v\n", err)
			}

			log.Printf("received: %v\n", res.Result)
		}
		close(waitChan)
	}()

	<-waitChan
}

func greetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("greetWithDeadline() invoked!")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Petros!"}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			// handle error
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadlline exceeded!")
			} else {
				log.Printf("unexpected gRPC err: %v\n", err)
			}
		} else {
			log.Printf("A non gRPC err: %v\n", err)
		}
	}

	log.Printf("greetWithDeadline: %s\n", res.Result)
}
