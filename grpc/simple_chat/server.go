package main

import (
	"fmt"
	"learn-microservices-go/grpc/simple_chat/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

// https://tutorialedge.net/golang/go-grpc-beginners-tutorial/#differences-between-grpc-and-rest

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
