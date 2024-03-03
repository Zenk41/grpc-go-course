package main

import (
	"context"
	"log"

	pb "github.com/Zenk41/grpc-go-course/greet/proto"
)


func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(),&pb.GreetRequest{
		FirstName: "Zenk41",
	})

	if err != nil{
		log.Fatalf("Could not greet: %v\n", err)

	}

	log.Printf("Greeting: %s\n", res.Result)
}