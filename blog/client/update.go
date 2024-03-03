package main

import (
	"context"
	"log"

	pb "github.com/Zenk41/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Not Zenk41",
		Title: "A new title",
		Content: "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}