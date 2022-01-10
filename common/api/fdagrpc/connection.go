package fdagrpc

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
)

const UrlPrefix = "/Api.Auth/"

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func GetGrpcConnection(authInterceptor grpc.UnaryClientInterceptor) (grpc.ClientConnInterface, error) {
	flag.Parse()
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(authInterceptor),
		grpc.WithTransportCredentials(creds),
	}
	return grpc.Dial(*addr, opts...)
}
