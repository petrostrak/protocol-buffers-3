package main

import (
	pb "blog-project/proto"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) CreateBLog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error)

func (s *Server) ListBlogs(_ *empty.Empty, stream pb.BlogService_ListBlogsServer) error
