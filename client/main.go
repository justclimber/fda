package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/justclimber/fda/common/api/generated/hello"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	res, err := client.SayHello(context.Background(), &pb.Name{Name: "Alex"})
	if err != nil {
		log.Fatalf("fail to call rpc say hello: %v", err)
	}
	if res.Success {
		log.Println("Success!")
	} else {
		log.Println("Something went wrong...")
	}
}
