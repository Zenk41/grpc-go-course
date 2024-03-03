package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Zenk41/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoke")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs :=[]*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}
	waitc := make(chan struct{})

	go func(){
		for _, req := range reqs {
			log.Printf("Send Request: %v\n", req)
			stream.Send(&pb.MaxRequest{
				Number: req.Number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func(){
		for {
			res, err := stream.Recv()

			if err == io.EOF{
				break
			}

			if err != nil {
				log.Printf("Problem while reading server stream: %v", err)
				break
			}

			log.Printf("Received a new maximum: %d\n", res.Result)
		}
		close(waitc)
	}()

<-waitc



}