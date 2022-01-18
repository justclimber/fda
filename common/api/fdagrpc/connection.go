package fdagrpc

import (
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const UrlPrefix = "/Api.Auth/"
const AuthKeyInMetadata = "authorization"

func GetGrpcConnection(addr string, authInterceptor grpc.UnaryClientInterceptor) (grpc.ClientConnInterface, error) {
	flag.Parse()
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(authInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	return grpc.Dial(addr, opts...)
}
