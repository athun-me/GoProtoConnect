package main

import (
	"log"
	"net"

	pb "github.com/athunlal/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type hellowServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &hellowServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Faild to start : %v", err)
	}
}
