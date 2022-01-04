package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/justclimber/fda/common/api/generated/api"
)

type User struct {
	id   uint64
	name string
}

func NewServer() *Server {
	return &Server{
		users:      make(map[uint64]User),
		lastUserId: 0,
	}
}

type Server struct {
	users      map[uint64]User
	lastUserId uint64

	grpcServer *grpc.Server
}

func (s *Server) Register(_ context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error) {
	s.lastUserId++
	id := s.lastUserId
	s.users[id] = User{id: id, name: in.Name}

	return &pb.RegisterOut{
		ID:      id,
		ErrCode: 0,
	}, nil
}

func (s *Server) Login(_ context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	user, found := s.users[in.ID]
	if !found {
		return &pb.LoginOut{
			ErrCode: 1,
		}, nil
	}

	return &pb.LoginOut{
		User: &pb.User{
			ID:   user.id,
			Name: user.name,
		},
		ErrCode: 0,
	}, nil
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s.grpcServer = grpc.NewServer(opts...)
	pb.RegisterAuthServer(s.grpcServer, s)
	log.Fatal(s.grpcServer.Serve(lis))
}
