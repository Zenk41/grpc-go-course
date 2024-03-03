package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Zenk41/grpc-go-course/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("prime was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling doPrime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)

		}
		log.Printf("Prime: %d\n", msg.Result)
	}
}


