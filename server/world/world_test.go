package world

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/server/player"
)

func TestLevelRegisterNewObj(t *testing.T) {
	l := NewWorld()
	_, p := player.NewPlayerWithComponent(3)
	e := NewPlayerEntity(123, p)
	err := l.RegisterNewEntity(e)
	require.NoError(t, err)
}
