package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/justclimber/fda/common/api/commonapi"
	"github.com/justclimber/fda/common/api/fdagrpc"
)

type AuthInterceptor struct {
	token string
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
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
	if !commonapi.AuthMethods[method] {
		return invoker(a.attachToken(ctx), method, req, reply, cc, opts...)
	}

	return invoker(ctx, method, req, reply, cc, opts...)
}

func (a *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, fdagrpc.AuthKeyInMetadata, a.token)
}
