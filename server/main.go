package main

import (
	"log"
	"net"

	"go-envoy.com/gen/services/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	grpcPort = "0.0.0.0:50051"
)

type Server struct{}

func (s *Server) SayHello(context context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	log.Printf("Received hello request!")

	if len(in.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Name not provided.")
	}

	return &helloworld.HelloResponse{Message: "Hello " + in.Name}, nil
}
func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	log.Printf("Listening at %s", grpcPort)
	grpcServer.Serve(listen)
}
