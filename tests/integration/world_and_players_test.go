package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/configloader"
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/debugger/templates"
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/player"
	"github.com/justclimber/fda/server/playersprocessor"
	"github.com/justclimber/fda/server/worldlog"
	"github.com/justclimber/fda/server/worldprocessor"
	"github.com/justclimber/fda/server/worldprocessor/ecs/generated/wprepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpsystem"
)

func TestWorldProcessorRun_WithPlayerProcessor(t *testing.T) {
	entityId := entity.Id(13)
	currentTick := tick.Tick(23)
	startPos := fgeom.Point{X: 8, Y: 20}
	objectivePos := startPos.Add(fgeom.Point{X: 2})
	tickLimit := tick.Tick(10)
	delay := 3
	sendLogsDelay := 4
	syncDelay := 2

	appCfg, err := configloader.NewConfigLoader().Load()
	require.NoError(t, err, "loading config")

	hr := debugger.NewHtmlReport(appCfg.DebuggerHtmlReportFullPath(), templates.EmbeddedFS, time.Second*2)
	d, finish := debugger.NewDebuggerWithReportFinish(true, hr)
	defer finish()

	wpDebugger := d.CreateNested("World Processor")
	ppDebugger := d.CreateNested("Players Processor")
	ecsDebugger := wpDebugger.CreateNested("ECS")
	playerCommandsDebugger := ecsDebugger.CreateNested("PlayerCommands")
	logDebugger := ecsDebugger.CreateNested("Log")

	ppWpLink := internalapi.NewPpWpLink()

	pl, plComp := player.NewPlayerWithComponent(delay)
	e := wprepo.EntityMask7{
		Id:       entityId,
		Position: wpcomponent.NewPosition(startPos),
		Moving:   wpcomponent.NewMoving(fgeom.EmptyPoint),
		Player:   plComp,
	}

	repo := entityrepo.NewChunked(wprepo.GetAllECGroups())

	repoForMask3 := wprepo.NewRepoForMask3(repo)
	repoForMask6 := wprepo.NewRepoForMask6(repo)

	l := worldlog.NewLogger()
	ec, err := ecs.NewEcs([]ecs.System{
		wpsystem.NewPlayerCommands(repoForMask6, delay, playerCommandsDebugger),
		wpsystem.NewMoving(repoForMask3),
		wpsystem.NewPosObjective(repoForMask3, entityId, objectivePos),
		wpsystem.NewTickLimiter(currentTick, tickLimit),
		wpsystem.NewLog(l, ppWpLink, sendLogsDelay, syncDelay, logDebugger),
	}, repo, ecsDebugger)
	require.NoError(t, err, "fail to create ecs")

	wp := worldprocessor.NewWorldProcessor(ec, ppWpLink, wpDebugger)
	require.NotNil(t, wp, "fail to create WorldProcessor")

	wp.AddEntity(e)

	pp := playersprocessor.NewPlayersProcessor(ppWpLink, ppDebugger)
	pp.AddPlayer(pl)

	go func() {
		err = pp.Run()
		require.NoError(t, err, "error while running PP")
	}()

	wp.Run(currentTick)
	require.NotEmpty(t, l.Logs(), "empty result logs")

	if l.Count() == int(tickLimit) {
		t.Fatal("tick limit reached unexpectedly")
	}

	expectedLogs := &worldlog.Logs{Entries: []worldlog.LogEntry{
		{Tick: 23},
		{Tick: 24},
		{Tick: 25},
		{Tick: 26},
		{Tick: 27},
		{Tick: 28},
	}}
	require.Equal(t, expectedLogs, l.Logs(), "check result logs")
}
