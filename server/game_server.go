package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/server/user"
)

type GameServer struct {
	authHelper *authHelper
}

func NewGameServer(usersFinder user.Finder) *GameServer {
	return &GameServer{authHelper: NewAuthHelper(usersFinder)}
}

func (g *GameServer) SomeMethodUnderAuth(ctx context.Context, empty *emptypb.Empty) (*pb.SomeRes, error) {
	u, err := g.authHelper.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.SomeRes{
		Success: true,
		Name:    u.Name,
	}, nil
}
