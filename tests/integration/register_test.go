//go:build integration

package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"

	"github.com/justclimber/fda/client"
	"github.com/justclimber/fda/common/api/commonapi"
	"github.com/justclimber/fda/common/api/fdagrpc"
	"github.com/justclimber/fda/common/configloader"
	"github.com/justclimber/fda/common/hasher/bcrypt"
	"github.com/justclimber/fda/server"
	"github.com/justclimber/fda/server/repository"
	"github.com/justclimber/fda/server/token"
)

type AuthClientServerSuit struct {
	suite.Suite
	s               *server.Server
	conn            grpc.ClientConnInterface
	cl              *client.AuthClient
	authInterceptor *client.AuthInterceptor
}

func TestAuthClientServerSuit(t *testing.T) {
	suite.Run(t, new(AuthClientServerSuit))
}

func (a *AuthClientServerSuit) SetupTest() {
	appCfg, err := configloader.NewConfigLoader().Load()
	require.NoError(a.T(), err)

	users := repository.NewUsersInMemory()
	hasher := bcrypt.Bcrypt{}
	tokenGenerator := new(token.SimpleTokenGenerator)

	a.s = server.NewServer(appCfg, users, hasher, tokenGenerator)
	go a.s.Start()

	a.authInterceptor = client.NewAuthInterceptor()
	a.conn, err = fdagrpc.GetGrpcConnection(appCfg.ServerUrl, a.authInterceptor.Unary)
	require.NoError(a.T(), err)

	a.cl = a.newAuthClient()
}

func (a *AuthClientServerSuit) TearDownTest() {
	a.s.Stop()
}

func (a *AuthClientServerSuit) newAuthClient() *client.AuthClient {
	cl, err := client.NewAuthClient(a.conn)
	require.NoError(a.T(), err)
	return cl
}

func (a *AuthClientServerSuit) registerOk(name string, password string) uint64 {
	res, err := a.cl.Register(name, password)
	require.NoError(a.T(), err)
	require.Equal(a.T(), uint32(0), res.ErrCode)
	return res.ID
}

func (a *AuthClientServerSuit) TestRegisterAndLogin_GetUserName() {
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
	for _, tc := range cases {
		a.T().Run(tc.caseName, func(t *testing.T) {
			id := a.registerOk(tc.userName, tc.password)
			res, err := a.cl.Login(id, tc.password)
			require.NoError(t, err)
			require.Equal(t, uint32(0), res.ErrCode)
			require.Equal(t, tc.userName, res.User.Name)
		})
	}
}

func (a *AuthClientServerSuit) TestRegisterDuplicate_GetLogicError() {
	const name = "Alex"
	a.registerOk(name, "123")

	res, err := a.cl.Register(name, "124")
	require.NoError(a.T(), err)
	require.Equal(a.T(), commonapi.RegisterUserAlreadyExists, res.ErrCode)
}

func (a *AuthClientServerSuit) TestRegisterWithEmptyName_GetLogicError() {
	cl := a.newAuthClient()

	res, err := cl.Register("", "123")
	require.NoError(a.T(), err)
	require.Equal(a.T(), commonapi.RegisterUserNameEmpty, res.ErrCode)
}

func (a *AuthClientServerSuit) TestRegisterAndLoginWithWrongPassword_GetLogicError() {
	id := a.registerOk("Alex", "right pass")

	res, err := a.cl.Login(id, "wrong pass")
	require.NoError(a.T(), err)
	require.Equal(a.T(), commonapi.LoginWrongPassword, res.ErrCode)
}

func (a *AuthClientServerSuit) TestLogin_ErrorNotFound() {
	const notExistedUserId = 987987987
	res, err := a.cl.Login(notExistedUserId, "")
	require.NoError(a.T(), err)
	require.Equal(a.T(), commonapi.LoginUserNotFound, res.ErrCode)
}
