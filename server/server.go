package server

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/justclimber/fda/common/api"
	"github.com/justclimber/fda/common/api/fdagrpc"
	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/common/hasher"
)

type Server struct {
	authServer *AuthServer
	gameServer *GameServer
	grpcServer *grpc.Server
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func NewServer(hasher hasher.Hasher) *Server {
	return &Server{
		authServer: NewAuthServer(hasher),
		gameServer: NewGameServer(),
	}
}

func (s *Server) Start() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s.grpcServer = grpc.NewServer(opts...)
	pb.RegisterAuthServer(s.grpcServer, s.authServer)
	pb.RegisterGameServer(s.grpcServer, s.gameServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == SecretToken
}

const SecretToken = "some-secret-token"

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	authMethods := map[string]bool{
		fdagrpc.UrlPrefix + "Register": true,
		fdagrpc.UrlPrefix + "Login":    true,
	}
	if authMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(md["authorization"]) {
		return nil, api.ErrUnauthorizedInvalidToken
	}
	return handler(ctx, req)
}
