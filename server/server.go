package server

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/justclimber/fda/common/api/commonapi"
	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/server/token"
	"github.com/justclimber/fda/server/user"
)

const ContextUserIdKey = "userId"
const authKeyInMetadata = "authorization"

type Server struct {
	authServer  *AuthServer
	gameServer  *GameServer
	grpcServer  *grpc.Server
	tokenFinder user.TokenFinder
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func NewServer(users user.Repository, hasher hasher.Hasher, tokenGenerator token.Generator) *Server {
	return &Server{
		authServer:  NewAuthServer(users, hasher, tokenGenerator),
		gameServer:  NewGameServer(users),
		tokenFinder: users,
	}
}

func (s *Server) Start() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(s.ensureValidToken),
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

func (s *Server) ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if commonapi.AuthMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	authorization := md[authKeyInMetadata]
	if len(authorization) == 0 {
		return nil, nil
	}

	u, err := s.tokenFinder.FindByToken(authorization[0])
	if u == nil || err != nil {
		return nil, commonapi.ErrUnauthorizedInvalidToken
	}
	return handler(context.WithValue(ctx, ContextUserIdKey, u.Id), req)
}
