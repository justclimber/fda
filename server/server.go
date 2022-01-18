package server

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/justclimber/fda/common/api/commonapi"
	"github.com/justclimber/fda/common/api/fdagrpc"
	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/common/config"
	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/server/token"
	"github.com/justclimber/fda/server/user"
)

const ContextUserIdKey = "userId"

type Server struct {
	cfg         config.Config
	authServer  *AuthServer
	gameServer  *GameServer
	grpcServer  *grpc.Server
	tokenFinder user.TokenFinder
}

func NewServer(cfg config.Config, users user.Repository, hasher hasher.Hasher, tokenGenerator token.Generator) *Server {
	return &Server{
		cfg:         cfg,
		authServer:  NewAuthServer(users, hasher, tokenGenerator),
		gameServer:  NewGameServer(users),
		tokenFinder: users,
	}
}

func (s *Server) Start() {
	flag.Parse()

	opts := []grpc.ServerOption{grpc.UnaryInterceptor(s.ensureValidToken)}

	s.grpcServer = grpc.NewServer(opts...)
	pb.RegisterAuthServer(s.grpcServer, s.authServer)
	pb.RegisterGameServer(s.grpcServer, s.gameServer)

	lis, err := net.Listen("tcp", s.cfg.ServerUrl)
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
		return nil, commonapi.ErrMissingMetadata
	}

	authorization := md[fdagrpc.AuthKeyInMetadata]
	if len(authorization) == 0 {
		return nil, nil
	}

	u, err := s.tokenFinder.FindByToken(authorization[0])
	if u == nil || err != nil {
		return nil, commonapi.ErrUnauthorizedInvalidToken
	}
	return handler(context.WithValue(ctx, ContextUserIdKey, u.Id), req)
}
