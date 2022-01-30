package level

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/server/player"
)

func TestLevelRegisterNewObj(t *testing.T) {
	l := NewLevel()
	p := player.NewPlayer()
	e := NewPlayerEntity(123, p)
	err := l.RegisterNewEntity(e)
	require.NoError(t, err)
}
