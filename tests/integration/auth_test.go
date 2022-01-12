package integration

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/client"
	"github.com/justclimber/fda/common/api/commonapi"
)

func (a *AuthClientServerSuit) TestAuth_Success() {
	password := "123"
	id := a.registerOk("Alex", password)

	lres, err := a.cl.Login(id, password)
	require.NoError(a.T(), err)
	require.Equal(a.T(), uint32(0), lres.ErrCode)

	a.authInterceptor.SetToken(lres.Token)
	gcl, err := client.NewGameClient(a.conn)
	gres, err := gcl.SomeMethodUnderAuth()
	require.NoError(a.T(), err)
	assert.True(a.T(), gres.Success)
}

func (a *AuthClientServerSuit) TestAuth_ErrorUnauthorized() {
	cl, err := client.NewGameClient(a.conn)
	_, err = cl.SomeMethodUnderAuth()

	assert.ErrorIs(a.T(), err, commonapi.ErrUnauthorizedInvalidToken)
}

func (a *AuthClientServerSuit) TestGetUserInfoThroughToken_Success() {
	password := "123"
	const name = "Alex"
	idAlex := a.registerOk(name, password)

	lres, err := a.cl.Login(idAlex, password)
	require.NoError(a.T(), err)
	require.Equal(a.T(), uint32(0), lres.ErrCode)

	a.authInterceptor.SetToken(lres.Token)
	gcl, err := client.NewGameClient(a.conn)
	gres, err := gcl.SomeMethodUnderAuth()
	require.NoError(a.T(), err)
	require.Equal(a.T(), name, gres.Name)
}
