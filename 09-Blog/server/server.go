package main

import (
	pb "blog-project/proto"
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBLog() invoked with %v\n", in)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal err: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error casting to oid: %v\n", err),
		)
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error)

func (s *Server) ListBlogs(_ *empty.Empty, stream pb.BlogService_ListBlogsServer) error
