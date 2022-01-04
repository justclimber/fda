package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/justclimber/fda/common/api/generated/api"
)

type Dialer interface {
	Dial(target string) (*grpc.ClientConn, error)
}

type Client struct {
	grpcClient api.AuthClient
}

func NewClient(dialer Dialer) (*Client, error) {
	conn, err := dialer.Dial("localhost:50051")
	if err != nil {
		return nil, err
	}
	return &Client{grpcClient: api.NewAuthClient(conn)}, nil
}

func (c *Client) Login(id uint64) (*api.LoginOut, error) {
	return c.grpcClient.Login(context.Background(), &api.LoginIn{ID: id})
}

func (c Client) Register(name string) (*api.RegisterOut, error) {
	return c.grpcClient.Register(context.Background(), &api.RegisterIn{
		Name: name,
	})
}
