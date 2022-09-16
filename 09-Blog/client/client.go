package main

import (
	pb "blog-project/proto"
	"context"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog() invoked!")

	blog := &pb.Blog{
		AuthorId: "Petros",
		Title:    "My First Blog",
		Content:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Printf("unexpected err %v\n", err)
	}

	log.Printf("Blog has been created with a new id of %s.\n", res.Id)

	return res.Id
}
