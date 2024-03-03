package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Zenk41/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("Avg was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doAvg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg : %v\n", err)
	}

	log.Printf("Avg: %.2f\n", res.Result)
}


