package main

import (
	"context"
	"log"

	pb "github.com/Zenk41/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with: %v\n", in)

	old, err := primitive.ObjectIDFromHex(in.Id)

	if err!= nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"cannot parse ID",
		)
	}

	data := &BlogItem{

	}

	filter := bson.M{"_id":old}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, "Cannot find blog withg ther ID provided",
		)
	}

	return documentToBlog(data), nil
}