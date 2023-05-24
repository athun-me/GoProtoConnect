package main

import (
	"log"
	"time"

	pb "github.com/athunlal/grpc-demo/proto"
)

func (s *hellowServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Fatalf("got requst with name : %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
