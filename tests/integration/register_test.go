package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	}{
		{
			caseName: "test Alex",
			userName: "Alex",
		},
		{
			caseName: "test Dunkan",
			userName: "Dunkan",
		},
	}
	s := server.NewServer()
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			cl, err := client.NewAuthClient(conn)
			assert.NoError(t, err)
			id := register(t, cl, tc.userName)
			login(t, cl, id, tc.userName)
		})
	}
}

func register(t *testing.T, cl *client.AuthClient, name string) uint64 {
	t.Helper()
	res, err := cl.Register(name)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), res.ErrCode)
	return res.ID
}

func login(t *testing.T, cl *client.AuthClient, id uint64, expectedName string) {
	t.Helper()
	res, err := cl.Login(id)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), res.ErrCode)
	assert.Equal(t, expectedName, res.User.Name)
}

func TestRegisterDuplicate_GetLogicError(t *testing.T) {
	s := server.NewServer()
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)

	cl, err := client.NewAuthClient(conn)
	require.NoError(t, err)

	const name = "Alex"
	register(t, cl, name)

	res, err := cl.Register(name)
	assert.NoError(t, err)
	assert.Equal(t, api.RegisterUserAlreadyExists, res.ErrCode)
}

func TestRegisterWithEmptyName_GetLogicError(t *testing.T) {
	s := server.NewServer()
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)
	cl, err := client.NewAuthClient(conn)
	assert.NoError(t, err)

	res, err := cl.Register("")
	assert.NoError(t, err)
	assert.Equal(t, api.RegisterUserNameEmpty, res.ErrCode)
}

func TestLogin_ErrorNotFound(t *testing.T) {
	s := server.NewServer()
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)
	cl, err := client.NewAuthClient(conn)
	assert.NoError(t, err)

	const notExistedUserId = 987987987
	res, err := cl.Login(notExistedUserId)
	assert.NoError(t, err)
	assert.Equal(t, api.LoginUserNotFound, res.ErrCode)
}
