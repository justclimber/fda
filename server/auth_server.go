package server

import (
	"context"
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/justclimber/fda/common/api"
	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/common/hasher/bcrypt"
	"github.com/justclimber/fda/server/token"
	"github.com/justclimber/fda/server/user"
)

func NewAuthServer(users user.Repository, hasher hasher.Hasher, tokenGenerator token.Generator) *AuthServer {
	return &AuthServer{
		users:          users,
		hasher:         hasher,
		tokenGenerator: tokenGenerator,
	}
}

type AuthServer struct {
	users          user.Repository
	hasher         hasher.Hasher
	tokenGenerator token.Generator
	grpcServer     *grpc.Server
}

var port = flag.Int("port", 50051, "the port to serve on")

func (a *AuthServer) Register(_ context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error) {
	userToRegister, err := user.NewUserToRegister(in.Name, in.Password, a.hasher)
	if err == user.ErrEmptyName {
		return &pb.RegisterOut{ErrCode: api.RegisterUserNameEmpty}, nil
	} else if err == user.ErrEmptyPassword {
		return &pb.RegisterOut{ErrCode: api.RegisterPasswordEmpty}, nil
	} else if err != nil {
		return nil, status.Error(codes.Internal, "can't composer user to register")
	}

	id, err := a.users.Register(userToRegister)
	if err == user.ErrAlreadyExists {
		return &pb.RegisterOut{ErrCode: api.RegisterUserAlreadyExists}, nil
	} else if err != nil {
		return nil, status.Error(codes.Internal, "can't register user")
	}

	return &pb.RegisterOut{
		ID:      id,
		ErrCode: 0,
	}, nil
}

func (a *AuthServer) Login(_ context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	u, err := a.users.FindById(in.ID)
	if u == nil {
		return &pb.LoginOut{ErrCode: api.LoginUserNotFound}, nil
	}

	check, err := u.CheckPassword(in.Password, bcrypt.Bcrypt{})
	if err != nil {
		return nil, status.Error(codes.Internal, "check password error")
	}

	if !check {
		return &pb.LoginOut{ErrCode: api.LoginWrongPassword}, nil
	}

	tok := u.GenerateToken(a.tokenGenerator)
	err = a.users.StoreToken(in.ID, tok)
	if err != nil {
		return nil, status.Error(codes.Internal, "can't store token")
	}

	return &pb.LoginOut{
		User: &pb.User{
			ID:   u.Id,
			Name: u.Name,
		},
		Token:   tok,
		ErrCode: 0,
	}, nil
}
