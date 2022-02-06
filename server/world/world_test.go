package world

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/server/player"
)

func TestWorldRegisterNewObj(t *testing.T) {
	w := NewWorld()
	_, p := player.NewPlayerWithComponent(3)
	e := entity.NewEntity(123)
	e.AddComponent(p)

	err := w.RegisterNewEntity(e)
	require.NoError(t, err)
}
