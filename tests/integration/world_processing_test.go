//go:build integration

package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/configloader"
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/debugger/templates"
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/player"
	"github.com/justclimber/fda/server/world"
	"github.com/justclimber/fda/server/worldlog"
	"github.com/justclimber/fda/server/worldprocessor"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpsystem"
)

func TestWorldProcessing_WithNewObj(t *testing.T) {
	w := world.NewWorld()
	_, pl := player.NewPlayerWithComponent(1)
	e := ecs.NewEntity(12)
	e.AddComponent(pl)

	err := w.RegisterNewEntity(e)
	require.NoError(t, err, "fail to register new object")

	err = w.AllocateEntity(e)
	require.NoError(t, err, "fail to allocate object")
}

func TestWorldProcessorRun_WithObjectiveAndTickLimiter(t *testing.T) {
	entityId := ecs.EntityId(13)
	currentTick := tick.Tick(23)
	startPosition := &fgeom.Point{X: 6, Y: 20}
	delay := 1
	sendLogsDelay := 1000

	appCfg, err := configloader.NewConfigLoader().Load()
	require.NoError(t, err, "loading config")

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

			hr := debugger.NewHtmlReport(appCfg.DebuggerHtmlReportFullPath(), templates.EmbeddedFS, time.Second*2)
			d, finish := debugger.NewDebuggerWithReportFinish(true, hr)
			defer finish()
			wpDebugger := d.CreateNested("World Processor")

			ppWpLink := internalapi.NewPpWpLink()

			ec, err := ecs.NewEcs([]ecs.System{
				wpsystem.NewMoving(),
				wpsystem.NewPosObjective(entityId, tc.posObjectivePoint),
				wpsystem.NewTickLimiter(currentTick, tc.tickLimit),
			}, wpDebugger.CreateNested("ECS"))
			require.NoError(t, err, "fail to create ecs")

			l := worldlog.NewLogger()
			wp := worldprocessor.NewWorldProcessor(l, ec, ppWpLink, sendLogsDelay, wpDebugger)
			require.NotNil(t, wp, "fail to create WorldProcessor")

			_, pl := player.NewPlayerWithComponent(delay)
			e := ecs.NewEntity(entityId)
			e.AddComponent(pl)

			e.AddComponent(&wpcomponent.Position{Pos: startPosition})
			e.AddComponent(wpcomponent.NewEngine(tc.power))

			err = wp.AddEntity(e)
			require.NoError(t, err, "fail to add entity")

			err = wp.Run(currentTick)
			require.NoError(t, err, "error while running WP&ECS")
			require.NotEmpty(t, l.Logs(), "empty result logs")
			require.Equal(t, tc.wantLogsCount, l.Count(), "check result logs count")
		})
	}
}
