//go:build integration

package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/ecs/servcomponent"
	"github.com/justclimber/fda/server/ecs/servsystem"
	"github.com/justclimber/fda/server/level"
	"github.com/justclimber/fda/server/levellog"
	"github.com/justclimber/fda/server/lpu"
	"github.com/justclimber/fda/server/player"
)

func TestLevelProcessing_WithNewObj(t *testing.T) {
	l := level.NewLevel()
	_, pl := player.NewPlayerWithComponent(1)
	e := level.NewPlayerEntity(12, pl)

	err := l.RegisterNewEntity(e)
	require.NoError(t, err, "fail to register new object")

	err = l.AllocateEntity(e)
	require.NoError(t, err, "fail to allocate object")
}

func TestLpuRun_WithObjectiveAndTickLimiter(t *testing.T) {
	entityId := ecs.EntityId(13)
	currentTick := tick.Tick(23)
	startPosition := &fgeom.Point{X: 6, Y: 20}

	for _, tc := range []struct {
		name              string
		posObjectivePoint fgeom.Point
		tickLimit         tick.Tick
		power             float64
		wantLogsCount     int
	}{
		{
			name:              "objective_reached_first",
			posObjectivePoint: startPosition.Add(fgeom.Point{X: 4}),
			tickLimit:         10,
			power:             1,
			wantLogsCount:     4,
		},
		{
			name:              "tick_limiter_reached_first",
			posObjectivePoint: startPosition.Add(fgeom.Point{X: 4}),
			tickLimit:         3,
			power:             1,
			wantLogsCount:     3,
		},
		{
			name:              "tick_limiter_reached_first_#2",
			posObjectivePoint: startPosition.Add(fgeom.Point{X: 10}),
			tickLimit:         4,
			power:             2,
			wantLogsCount:     4,
		},
		{
			name:              "tick_limiter_with_objective",
			posObjectivePoint: startPosition.Add(fgeom.Point{X: 4}),
			tickLimit:         4,
			power:             1,
			wantLogsCount:     4,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			moving := servsystem.NewMoving()
			posObjective := servsystem.NewPosObjective(entityId, tc.posObjectivePoint)
			tickLimiter := servsystem.NewTickLimiter(currentTick, tc.tickLimit)
			ec, err := ecs.NewEcs([]ecs.System{moving, posObjective, tickLimiter})
			require.NoError(t, err, "fail to create ecs")

			log := levellog.NewLevelLog()
			lp := lpu.NewLpu(log, ec)
			require.NotNil(t, lp, "fail to create LPU")

			_, pl := player.NewPlayerWithComponent(1)
			e := level.NewPlayerEntity(entityId, pl)

			e.AddComponent(servcomponent.CPosition, &servcomponent.Position{Pos: startPosition})
			e.AddComponent(servcomponent.CMovable, servcomponent.NewEngine(tc.power))

			err = lp.AddEntity(e)
			require.NoError(t, err, "fail to add entity")

			err = lp.Run(currentTick)
			require.NoError(t, err, "error while running LPU&ECS")
			require.NotEmpty(t, log.Logs(), "empty result logs")
			require.Len(t, log.Logs(), tc.wantLogsCount, "check result logs count")
		})
	}
}

func TestLpuRun_WithPpu(t *testing.T) {
	entityId := ecs.EntityId(13)
	currentTick := tick.Tick(23)
	startPos := &fgeom.Point{X: 8, Y: 20}
	objectivePos := startPos.Add(fgeom.Point{X: 2})
	power := 1.
	tickLimit := tick.Tick(10)

	moving := servsystem.NewMoving()
	posObjective := servsystem.NewPosObjective(entityId, objectivePos)
	playerCommands := servsystem.NewPlayerCommands()
	tickLimiter := servsystem.NewTickLimiter(currentTick, tickLimit)
	ec, err := ecs.NewEcs([]ecs.System{
		playerCommands,
		moving,
		posObjective,
		tickLimiter,
	})
	require.NoError(t, err, "fail to create ecs")

	log := levellog.NewLevelLog()
	lp := lpu.NewLpu(log, ec)
	require.NotNil(t, lp, "fail to create LPU")

	pl, plComp := player.NewPlayerWithComponent(3)
	e := level.NewPlayerEntity(entityId, plComp)
	e.AddComponent(servcomponent.CPosition, &servcomponent.Position{Pos: startPos})
	e.AddComponent(servcomponent.CMovable, servcomponent.NewEngine(power))

	err = lp.AddEntity(e)
	require.NoError(t, err, "fail to add entity")

	pl.SendCommand(command.Command{Move: 0.5})

	err = lp.Run(currentTick)
	require.NoError(t, err, "error while running LPU&ECS")
	require.NotEmpty(t, log.Logs(), "empty result logs")

	if len(log.Logs()) == int(tickLimit) {
		t.Fatal("tick limit reached unexpectedly")
	}

	expectedLogs := []*levellog.LogEntry{
		{Tick: 23},
		{Tick: 24},
		{Tick: 25},
		{Tick: 26},
	}
	require.Equal(t, expectedLogs, log.Logs(), "check result logs")
}
