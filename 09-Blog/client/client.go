package main

import (
	pb "blog-project/proto"
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog() invoked!")

	blog := &pb.Blog{
		AuthorId: "Eirini",
		Title:    "Vietnam Blog",
		Content:  "Vietnam is a Southeast Asian country known for its beaches, rivers, Buddhist pagodas and bustling cities. Hanoi, the capital, pays homage to the nation’s iconic Communist-era leader, Ho Chi Minh, via a huge marble mausoleum. Ho Chi Minh City (formerly Saigon) has French colonial landmarks, plus Vietnamese War history museums and the Củ Chi tunnels, used by Viet Cong soldiers.",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Printf("unexpected err %v\n", err)
	}

	log.Printf("Blog has been created with a new id of %s.\n", res.Id)

	return res.Id
}

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog() invoked!")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("err while reading %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog() invoked!")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Petros",
		Title:    "An updated Title",
		Content:  "Vietnam was awesome!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Printf("err while updating %v\n", err)
	}

	log.Println("Blog was updated!")
}

func listBlog(c pb.BlogServiceClient) {
	log.Println("listBlog() invoked!")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Error while calling ListBlog: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Error while receiving from ListBlog: %v\n", err)
		}

		log.Println(res)
	}
}
