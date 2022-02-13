package wpsystem

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/worldlog"
	"github.com/justclimber/fda/server/worldprocessor/ecs/generated/wprepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type Log struct {
	entityRepo    *wprepo.RepoForMask3
	logger        *worldlog.Logger
	ppWpLink      *internalapi.PpWpLink
	logsIndex     int
	sendLogsDelay int
	syncDelay     int
	debugger      *debugger.Nested
	LogBatch      worldlog.LogBatch
}

func NewLog(
	entityRepo *wprepo.RepoForMask3,
	logger *worldlog.Logger,
	ppWpLink *internalapi.PpWpLink,
	sendLogsDelay int,
	syncDelay int,
	debugger *debugger.Nested,
) *Log {
	return &Log{
		entityRepo:    entityRepo,
		logger:        logger,
		ppWpLink:      ppWpLink,
		sendLogsDelay: sendLogsDelay,
		syncDelay:     syncDelay,
		debugger:      debugger,
	}
}

func (l *Log) String() string { return "Log" }

func (l *Log) Init(tick tick.Tick) {
	l.LogBatch = worldlog.NewLogBatch(tick)
	l.LogBatch.EndTick = tick
	l.sendLog(tick - 1)
}

func (l *Log) DoTick(tick tick.Tick) bool {
	l.entityRepo.Iterate(func(
		id entity.Id,
		p wpcomponent.Position,
		m wpcomponent.Moving,
	) (*wpcomponent.Position, *wpcomponent.Moving) {
		l.LogBatch.Add(tick, id, m)

		return nil, nil
	})
	l.logger.LogTick(tick)

	l.sendLogAndSync(tick)
	return false
}

func (l *Log) sendLogAndSync(t tick.Tick) {
	l.logsIndex++
	if l.logsIndex >= l.sendLogsDelay {
		l.sendLog(t)
		l.logsIndex = 0
	} else if l.logsIndex == l.syncDelay {
		l.debugger.LogF("DoTick", "wait sync")
		<-l.ppWpLink.SyncCh
		l.debugger.LogF("DoTick", "sync get")
	}
}

func (l *Log) sendLog(t tick.Tick) {
	l.debugger.LogF("DoTick", "send logs")
	l.ppWpLink.LogsCh <- l.LogBatch
	l.logger.LogBatch(l.LogBatch)
	l.LogBatch = worldlog.NewLogBatch(t + 1)
}
