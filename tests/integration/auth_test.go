package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/client"
	"github.com/justclimber/fda/common/api/grpc"
	"github.com/justclimber/fda/common/hasher/bcrypt"
	"github.com/justclimber/fda/server"
)

var bcryptHasher = bcrypt.Bcrypt{}

func TestAuth_Success(t *testing.T) {
	s := server.NewServer(bcryptHasher)
	go s.Start()
	defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)

	cl, err := client.NewAuthClient(conn)
	require.NoError(t, err)
	res, err := cl.Register("Alex", "")
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.ErrCode)

	lres, err := cl.Login(res.ID, "")
	require.NoError(t, err)
	require.Equal(t, uint32(0), lres.ErrCode)

	gcl, err := client.NewGameClient(conn)
	gres, err := gcl.SomeMethodUnderAuth()
	require.NoError(t, err)
	assert.True(t, gres.Success)
}

func TestAuth_ErrorUnauthorized(t *testing.T) {
	s := server.NewServer(bcryptHasher)
	go s.Start()
	//defer s.Stop()

	conn, err := grpc.GetGrpcConnection()
	require.NoError(t, err)

	cl, err := client.NewGameClient(conn)
	_, err = cl.SomeMethodUnderAuth()
	assert.Error(t, err)
}
