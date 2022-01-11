package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/server/user"
)

type GameServer struct {
	usersFinder user.Finder
}

func NewGameServer(usersFinder user.Finder) *GameServer {
	return &GameServer{usersFinder: usersFinder}
}

func (g *GameServer) SomeMethodUnderAuth(ctx context.Context, empty *emptypb.Empty) (*pb.SomeRes, error) {
	id, ok := ctx.Value(ContextUserIdKey).(uint64)
	if !ok {
		return &pb.SomeRes{}, fmt.Errorf("empty auth token")
	}
	u, err := g.usersFinder.FindById(id)
	if err != nil {
		return &pb.SomeRes{}, fmt.Errorf("can't get user with auth token")
	}

	return &pb.SomeRes{
		Success: true,
		Name:    u.Name,
	}, nil
}
