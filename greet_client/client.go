package main

import (
	"context"
	"fmt"
	"github.com/raphaelluethy/go-go-brr/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	r, err := c.GetGreeting(context.Background(), &pb.GreetRequest{Name: "Peter"})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	fmt.Printf("Greeting: %s\n", r.Greeting)
}
