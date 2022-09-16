package main

import (
	pb "blog-project/proto"
	"context"
	"log"
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
		log.Printf("err while reading err %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
