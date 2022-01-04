package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/justclimber/fda/common/api/generated/api"
)

type HelloService struct {
}

func (h HelloService) SayHello(ctx context.Context, name *pb.Name) (*pb.Result, error) {
	log.Println("log: ", name.Name)
	return &pb.Result{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServer(grpcServer, &HelloService{})
	_ = grpcServer.Serve(lis)
}
