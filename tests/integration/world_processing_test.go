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
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/player"
	"github.com/justclimber/fda/server/worldlog"
	"github.com/justclimber/fda/server/worldprocessor"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wprepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpsystem"
)

func TestWorldProcessorRun_WithObjectiveAndTickLimiter(t *testing.T) {
	entityId := entity.Id(13)
	currentTick := tick.Tick(23)
	startPosition := fgeom.Point{X: 6, Y: 20}
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

			_, pl := player.NewPlayerWithComponent(delay)
			e := entity.NewEntity(entityId)
			e.AddComponent(pl)
			e.AddComponent(wpcomponent.NewPosition(startPosition))
			e.AddComponent(wpcomponent.NewMoving(fgeom.Point{X: tc.power}))

			cg7 := wprepo.NewCGroup7()
			repo := entityrepo.NewChunked(map[component.Mask]entityrepo.CGroup{
				e.CMask: cg7,
			})

			repoForMask3 := wprepo.NewRepoForMask3(repo)

			l := worldlog.NewLogger()
			ecsDebugger := wpDebugger.CreateNested("ECS")
			ec, err := ecs.NewEcs([]ecs.System{
				wpsystem.NewMoving(repoForMask3),
				wpsystem.NewPosObjective(repoForMask3, entityId, tc.posObjectivePoint),
				wpsystem.NewTickLimiter(currentTick, tc.tickLimit),
				wpsystem.NewLog(l, ppWpLink, sendLogsDelay, sendLogsDelay-2, ecsDebugger.CreateNested("Log")),
			}, repo, ecsDebugger)
			require.NoError(t, err, "fail to create ecs")

			wp := worldprocessor.NewWorldProcessor(ec, ppWpLink, wpDebugger)
			require.NotNil(t, wp, "fail to create WorldProcessor")

			wp.AddEntity(e)
			wp.Run(currentTick)
			require.NotEmpty(t, l.Logs(), "empty result logs")
			require.Equal(t, tc.wantLogsCount, l.Count(), "check result logs count")
		})
	}
}
