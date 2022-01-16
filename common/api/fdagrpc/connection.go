package fdagrpc

import (
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const UrlPrefix = "/Api.Auth/"

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func GetGrpcConnection(authInterceptor grpc.UnaryClientInterceptor) (grpc.ClientConnInterface, error) {
	flag.Parse()
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(authInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	return grpc.Dial(*addr, opts...)
}
