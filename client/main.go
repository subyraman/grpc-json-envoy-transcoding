package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	helloworld "go-envoy/gen/services/helloworld"

	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

func main() {
	address := "0.0.0.0:50051"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	// Set up a connection to the server.
	fmt.Printf("connecting to %s", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("attempting to say hello")
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}