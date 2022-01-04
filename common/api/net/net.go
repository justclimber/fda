package net

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcDialer struct{}

func (g GrpcDialer) Dial(target string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.Dial(target, opts...)
}
