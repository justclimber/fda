package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/justclimber/fda/common/api/fdagrpc"
	"github.com/justclimber/fda/common/api/generated/api"
)

type AuthClient struct {
	grpcClient api.AuthClient
	token      string
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

type AuthInterceptor struct {
	authMethods map[string]bool
	token       string
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{
		authMethods: map[string]bool{
			fdagrpc.UrlPrefix + "Register": true,
			fdagrpc.UrlPrefix + "Login":    true,
		},
	}
}

func (a *AuthInterceptor) SetToken(t string) {
	a.token = t
}

func (a *AuthInterceptor) Unary(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	if !a.authMethods[method] {
		return invoker(a.attachToken(ctx), method, req, reply, cc, opts...)
	}

	return invoker(ctx, method, req, reply, cc, opts...)
}

func (a *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", a.token)
}
