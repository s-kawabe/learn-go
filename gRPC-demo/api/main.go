package main

import (
	"context"
	"log"
	"net"

	pb "github.com/s-kawabe/learn-go/gRPC-demo/api/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := r.GetName()
	email := r.GetEmail()
	age := r.GetAge()
	log.Printf("Received : %s %d %s", name, email, age)
	// return response linked take message
	return &pb.HelloResponse{Message: "Hello, %s %d! your email is %s", name, age, email}, nil
}

func main() {
	// setting listen post
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// var opts []grpc.ServerOption

	// convert instance to server
	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)

	// test output log for runtime
	log.Println("Waiting for gRPC request ....")
	log.Println("--------------")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %c", err)
	}
}
