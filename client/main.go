package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/athunlal/grpc-demo/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	// client := pb.NewGreetServiceClient(conn)

	// name := &pb.NameList{
	// 	Names: []string{"Athun", "Arjun", "Bob"},
	// }

	// pb.Say
}
