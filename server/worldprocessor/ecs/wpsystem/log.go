package wpsystem

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs/component"
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

func (l *Log) Init(_ tick.Tick) {
	l.sendLog()
}

func (l *Log) DoTick(tick tick.Tick) bool {
	var needSend, needSync, needSnapshot bool
	var eComps map[entity.Id]map[component.Key]component.Component

	l.logsIndex++
	if l.logsIndex == 1 {
		needSnapshot = true
		eComps = map[entity.Id]map[component.Key]component.Component{}
	} else if l.logsIndex >= l.sendLogsDelay {
		needSend = true
	} else if l.logsIndex == l.syncDelay {
		needSync = true
	}

	l.entityRepo.Iterate(func(
		id entity.Id,
		p wpcomponent.Position,
		m wpcomponent.Moving,
	) (*wpcomponent.Position, *wpcomponent.Moving) {
		l.logger.AddToCurBatch(tick, id, m)
		if needSnapshot {
			eComps[id] = map[component.Key]component.Component{p.Key(): p}
		}
		return nil, nil
	})
	l.logger.LogTick(tick, eComps)

	if needSend {
		l.sendLog()
		l.logsIndex = 0
	} else if needSync {
		l.debugger.LogF("DoTick", "wait sync")
		<-l.ppWpLink.SyncCh
		l.debugger.LogF("DoTick", "sync get")
	}

	return false
}

func (l *Log) sendLog() {
	l.debugger.LogF("DoTick", "send logs")
	l.ppWpLink.LogsCh <- l.logger.RotateBatch()
}
