package main

import (
	"context"
	"log"
	"os"
	"time"

	helloworld "go-envoy.com/gen/services/helloworld"

	"google.golang.org/grpc"
)

const (
	defaultName = ""
)

func main() {
	address := "0.0.0.0:50051"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	// Set up a connection to the server.
	log.Printf("connecting to %s", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := ""
	if len(os.Args) > 2 {
		name = os.Args[2]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Received greeting from server: %s", r.GetMessage())
}
