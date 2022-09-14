package main

import (
	"log"
	"time"

	pb "07-grpc-greet-project/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr = "0.0.0.0:50051"
)

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Printf("error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Printf("failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// greet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	greetWithDeadline(c, 1*time.Second)
}
