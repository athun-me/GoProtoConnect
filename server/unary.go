package main

import (
	"context"

	pb "github.com/athunlal/grpc-demo/proto"
)

func (s *hellowServer) SayHello(ctx context.Context, in *pb.NoParm) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
