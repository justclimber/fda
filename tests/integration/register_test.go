package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/client"
	"github.com/justclimber/fda/common/api"
	"github.com/justclimber/fda/common/api/grpc"
	"github.com/justclimber/fda/server"
)

func TestRegisterAndLogin_GetUserName(t *testing.T) {
	cases := []struct {
		caseName string
		userName string
		password string
	}{
		{
			caseName: "test Alex",
			userName: "Alex",
			password: "pswd1",
		},
		{
			caseName: "test Dunkan",
			userName: "Dunkan",
			password: "pswd2",
		},
	}
	s := server.NewServer(bcryptHasher)
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			cl, err := client.NewAuthClient(conn)
			require.NoError(t, err)
			id := register(t, cl, tc.userName, tc.password)
			login(t, cl, id, tc.userName, tc.password)
		})
	}
}

func register(t *testing.T, cl *client.AuthClient, name string, password string) uint64 {
	t.Helper()
	res, err := cl.Register(name, password)
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.ErrCode)
	return res.ID
}

func login(t *testing.T, cl *client.AuthClient, id uint64, expectedName string, password string) {
	t.Helper()
	res, err := cl.Login(id, password)
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.ErrCode)
	require.Equal(t, expectedName, res.User.Name)
}

func TestRegisterDuplicate_GetLogicError(t *testing.T) {
	s := server.NewServer(bcryptHasher)
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)

	cl, err := client.NewAuthClient(conn)
	require.NoError(t, err)

	const name = "Alex"
	register(t, cl, name, "123")

	res, err := cl.Register(name, "124")
	require.NoError(t, err)
	require.Equal(t, api.RegisterUserAlreadyExists, res.ErrCode)
}

func TestRegisterWithEmptyName_GetLogicError(t *testing.T) {
	s := server.NewServer(bcryptHasher)
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)
	cl, err := client.NewAuthClient(conn)
	require.NoError(t, err)

	res, err := cl.Register("", "")
	require.NoError(t, err)
	require.Equal(t, api.RegisterUserNameEmpty, res.ErrCode)
}

func TestRegisterAndLoginWithWrongPassword_GetLogicError(t *testing.T) {
	s := server.NewServer(bcryptHasher)
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)
	cl, err := client.NewAuthClient(conn)
	require.NoError(t, err)

	id := register(t, cl, "Alex", "right pass")

	res, err := cl.Login(id, "wrong pass")
	require.NoError(t, err)
	require.Equal(t, api.LoginWrongPassword, res.ErrCode)
}

func TestLogin_ErrorNotFound(t *testing.T) {
	s := server.NewServer(bcryptHasher)
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)
	cl, err := client.NewAuthClient(conn)
	require.NoError(t, err)

	const notExistedUserId = 987987987
	res, err := cl.Login(notExistedUserId, "")
	require.NoError(t, err)
	require.Equal(t, api.LoginUserNotFound, res.ErrCode)
}
