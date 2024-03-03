package main

import (
	"log"

	pb "github.com/Zenk41/grpc-go-course/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error{
	log.Printf("Prime function was invoked with: %v\n", in)
	var k int32 = 2
	var n int32 = in.Number
	for n > 1 {
		if n %2 == 0 {
			stream.Send(&pb.PrimeResponse{
				Result:k,
			})
			n /=k
		} else {
			k +=1
		}
	}
	return nil
}