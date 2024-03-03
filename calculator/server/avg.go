package main

import (
	"io"
	"log"

	pb "github.com/Zenk41/grpc-go-course/calculator/proto"
)

func (server *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg Function has been invoked")

	var counter float32 = 0
	var sum float32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Avg: %v\n", sum/counter)
			return stream.SendAndClose(
				&pb.AvgResponse{
					Result: sum/counter,
				},
			)
		}
		if err != nil{
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving: %v\n", req)

		sum+=req.Number
		counter+=1
	}
	
}
