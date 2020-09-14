package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// https://tutorialedge.net/golang/go-grpc-beginners-tutorial/#differences-between-grpc-and-rest

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
