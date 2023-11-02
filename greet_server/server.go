package main

import (
	"context"
	"github.com/raphaelluethy/go-go-brr/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s server) GetGreeting(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Greeting: "Hello World!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
