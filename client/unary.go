package main

import (
	"context"
	"log"
	"time"

	pb "github.com/athunlal/grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParm{})
	if err != nil {
		log.Fatalf("could not greet : %v", err)
	}
	log.Printf("%s", res.Message)

}
