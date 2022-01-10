package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/justclimber/fda/common/api/generated/api"
)

type GameServer struct{}

func NewGameServer() *GameServer {
	return &GameServer{}
}

func (g *GameServer) SomeMethodUnderAuth(ctx context.Context, empty *emptypb.Empty) (*pb.Result, error) {
	return &pb.Result{
		Success: true,
	}, nil
}
