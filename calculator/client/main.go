package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/Zenk41/grpc-go-course/calculator/proto"
)

var addr string = "localhost:5001"

func main() {
	conn, err := grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()


	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrime(c)
	// doAvg(c)
	// doMax(c)
	doSqrt(c, -2)
}