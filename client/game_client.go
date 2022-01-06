package client

import (
	"google.golang.org/grpc"

	"github.com/justclimber/fda/common/api/generated/api"
)

type GameClient struct {
	grpcClient api.GameClient
}

func NewGameClient(c grpc.ClientConnInterface) (*GameClient, error) {
	return &GameClient{grpcClient: api.NewGameClient(c)}, nil
}

func (c *GameClient) SomeMethodUnderAuth() (*api.Result, error) {
	return &api.Result{
		Success: true,
	}, nil
}
