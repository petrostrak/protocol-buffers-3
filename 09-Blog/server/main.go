package main

import (
	pb "blog-project/proto"
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr       = "0.0.0.0:50051"
	collection *mongo.Collection
)

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Printf("failed to create new mongoDB client: %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Printf("failed to connect to to mongoDB: %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s...\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)

	err = s.Serve(l)
	if err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
