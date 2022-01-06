package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/justclimber/fda/common/api/generated/api"
)

type AuthClient struct {
	grpcClient api.AuthClient
}

func NewAuthClient(c grpc.ClientConnInterface) (*AuthClient, error) {
	return &AuthClient{grpcClient: api.NewAuthClient(c)}, nil
}

func (c *AuthClient) Login(id uint64, password string) (*api.LoginOut, error) {
	return c.grpcClient.Login(context.Background(), &api.LoginIn{
		ID:       id,
		Password: password,
	})
}

func (c AuthClient) Register(name string, password string) (*api.RegisterOut, error) {
	return c.grpcClient.Register(context.Background(), &api.RegisterIn{
		Name:     name,
		Password: password,
	})
}
