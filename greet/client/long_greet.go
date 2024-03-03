package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Zenk41/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Zenk41"},
		{FirstName: "Dhi"},
		{FirstName: "Wow"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreetL %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet : %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
