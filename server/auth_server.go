package server

import (
	"context"
	"flag"

	"google.golang.org/grpc"

	"github.com/justclimber/fda/common/api"
	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/common/hasher/bcrypt"
	"github.com/justclimber/fda/server/user"
)

func NewAuthServer(hasher hasher.Hasher) *AuthServer {
	return &AuthServer{
		users:         make(map[uint64]*user.User),
		userIdsByName: make(map[string]uint64),
		lastUserId:    0,
		hasher:        hasher,
	}
}

type AuthServer struct {
	users         map[uint64]*user.User
	userIdsByName map[string]uint64
	lastUserId    uint64

	hasher     hasher.Hasher
	grpcServer *grpc.Server
}

var port = flag.Int("port", 50051, "the port to serve on")

func (a *AuthServer) Register(_ context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error) {
	if in.Name == "" {
		return &pb.RegisterOut{ErrCode: api.RegisterUserNameEmpty}, nil
	}
	if _, found := a.userIdsByName[in.Name]; found {
		return &pb.RegisterOut{ErrCode: api.RegisterUserAlreadyExists}, nil
	}

	a.lastUserId++
	id := a.lastUserId
	newUser, err := user.NewUser(id, in.Name, in.Password, a.hasher)

	if err == user.EmptyPasswordProvided {
		return &pb.RegisterOut{ErrCode: api.RegisterPasswordEmpty}, nil
	} else if err != nil {
		return &pb.RegisterOut{ErrCode: api.InternalError}, err
	}

	a.users[id] = newUser
	a.userIdsByName[in.Name] = id

	return &pb.RegisterOut{
		ID:      id,
		ErrCode: 0,
	}, nil
}

func (a *AuthServer) Login(_ context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	u, found := a.users[in.ID]
	if !found {
		return &pb.LoginOut{ErrCode: api.LoginUserNotFound}, nil
	}

	check, err := u.CheckPassword(in.Password, bcrypt.Bcrypt{})
	if err != nil {
		return &pb.LoginOut{ErrCode: api.InternalError}, err
	}

	if !check {
		return &pb.LoginOut{ErrCode: api.LoginWrongPassword}, nil
	}

	return &pb.LoginOut{
		User: &pb.User{
			ID:   u.Id,
			Name: u.Name,
		},
		Token:   SecretToken,
		ErrCode: 0,
	}, nil
}
