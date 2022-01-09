package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/justclimber/fda/common/api/generated/api"
)

type GameClient struct {
	grpcClient api.GameClient
}

func NewGameClient(c grpc.ClientConnInterface) (*GameClient, error) {
	return &GameClient{grpcClient: api.NewGameClient(c)}, nil
}

func (c *GameClient) SomeMethodUnderAuth() (*api.Result, error) {
	return c.grpcClient.SomeMethodUnderAuth(context.Background(), &emptypb.Empty{})
}
