package server

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/justclimber/fda/common/api"
	pb "github.com/justclimber/fda/common/api/generated/api"
	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/common/hasher/bcrypt"
	"github.com/justclimber/fda/server/user"
)

func NewServer(hasher hasher.Hasher) *Server {
	return &Server{
		users:         make(map[uint64]*user.User),
		userIdsByName: make(map[string]uint64),
		lastUserId:    0,
		hasher:        hasher,
	}
}

type Server struct {
	users         map[uint64]*user.User
	userIdsByName map[string]uint64
	lastUserId    uint64

	hasher     hasher.Hasher
	grpcServer *grpc.Server
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

var port = flag.Int("port", 50051, "the port to serve on")

func (s *Server) Register(_ context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error) {
	if in.Name == "" {
		return &pb.RegisterOut{ErrCode: api.RegisterUserNameEmpty}, nil
	}
	if _, found := s.userIdsByName[in.Name]; found {
		return &pb.RegisterOut{ErrCode: api.RegisterUserAlreadyExists}, nil
	}

	s.lastUserId++
	id := s.lastUserId
	newUser, err := user.NewUser(id, in.Name, in.Password, s.hasher)

	if err == user.EmptyPasswordProvided {
		return &pb.RegisterOut{ErrCode: api.RegisterPasswordEmpty}, nil
	} else if err != nil {
		return &pb.RegisterOut{ErrCode: api.InternalError}, err
	}

	s.users[id] = newUser
	s.userIdsByName[in.Name] = id

	return &pb.RegisterOut{
		ID:      id,
		ErrCode: 0,
	}, nil
}

func (s *Server) Login(_ context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	u, found := s.users[in.ID]
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
		ErrCode: 0,
	}, nil
}

func (s *Server) Start() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s.grpcServer = grpc.NewServer(opts...)
	pb.RegisterAuthServer(s.grpcServer, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
