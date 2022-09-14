package main

import (
	pb "07-grpc-greet-project/greet/proto"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Unary API Server Implementation
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

// Server Streaming Implementation
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes() invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}

// Client Streaming Server Implementation
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("GreetManyTimes() invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		} else if err != nil {
			log.Printf("error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone() invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Printf("error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Printf("error while sending data to client: %v\n", err)
		}
	}
}

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline() invoked with: %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}

		time.Sleep(time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName + " !",
	}, nil
}
