package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/ecs/servcomponent"
	"github.com/justclimber/fda/server/ecs/servsystem"
	"github.com/justclimber/fda/server/level"
	"github.com/justclimber/fda/server/levellog"
	"github.com/justclimber/fda/server/lpu"
	"github.com/justclimber/fda/server/player"
)

func TestLevelProcessing_WithNewObj(t *testing.T) {
	l := level.NewLevel()
	pl := player.NewPlayer()
	e := level.NewPlayerEntity(12, pl)

	err := l.RegisterNewEntity(e)
	require.NoError(t, err, "fail to register new object")

	err = l.AllocateEntity(e)
	require.NoError(t, err, "fail to allocate object")
}

func TestLpuRun_WithObjective(t *testing.T) {
	entityId := ecs.EntityId(13)

	log := levellog.NewLevelLog()
	moving := servsystem.NewMoving()
	posObjective := servsystem.NewPosObjective(entityId, fgeom.Point{
		X: 10,
		Y: 20,
	})
	ec, err := ecs.NewEcs([]ecs.System{moving, posObjective})
	require.NoError(t, err)

	lp := lpu.NewLpu(log, ec)
	require.NotNil(t, lp)

	pl := player.NewPlayer()
	e := level.NewPlayerEntity(entityId, pl)
	e.AddComponent(servcomponent.CPosition, &servcomponent.Position{Pos: &fgeom.Point{X: 6, Y: 2}})
	e.AddComponent(servcomponent.CMovable, servcomponent.NewEngine(1))

	err = lp.AddEntity(e)
	require.NoError(t, err)

	currentTick := tick.Tick(23)
	err = lp.Run(currentTick)
	require.NoError(t, err)

	resultLogs := lp.Logger().Logs()
	require.NotEmpty(t, resultLogs)

	expectedLogs := []*levellog.LogEntry{
		{Tick: 23},
		{Tick: 24},
		{Tick: 25},
		{Tick: 26},
	}
	require.Equal(t, expectedLogs, resultLogs)
}
