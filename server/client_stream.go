package main

import (
	"io"
	"log"

	pb "github.com/athunlal/grpc-demo/proto"
)

func (s *hellowServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var message []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Message: message})
		}

		if err != nil {
			return err
		}
		log.Fatalf("Got requst with name : %v", req.Name)

		message = append(message, "Hello", req.Name)
	}
}
