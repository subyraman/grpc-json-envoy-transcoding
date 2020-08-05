package main

import (
	"fmt"
	"log"
	"net"

	"go-envoy/gen/services/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = "0.0.0.0:50051"
)

type Server struct{}

func (s *Server) SayHello(context context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	log.Printf("Received hello request!")
	return &helloworld.HelloResponse{Message: "Hello " + in.Name}, nil
}
func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	fmt.Printf("Listening at %s", grpcPort)
	grpcServer.Serve(listen)
}
