package wpsystem

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/worldlog"
)

type Log struct {
	logger        *worldlog.Logger
	ppWpLink      *internalapi.PpWpLink
	logsIndex     int
	sendLogsDelay int
	syncDelay     int
	debugger      *debugger.Nested
}

func NewLog(
	logger *worldlog.Logger,
	ppWpLink *internalapi.PpWpLink,
	sendLogsDelay int,
	syncDelay int,
	debugger *debugger.Nested,
) *Log {
	return &Log{
		logger:        logger,
		ppWpLink:      ppWpLink,
		sendLogsDelay: sendLogsDelay,
		syncDelay:     syncDelay,
		debugger:      debugger,
	}
}

func (l *Log) String() string {
	return "Log"
}

func (l *Log) Init() {
	l.sendLog()
}

func (l *Log) RequiredComponentKeys() []component.Key {
	return []component.Key{}
}

func (l *Log) AddEntity(_ *entity.Entity, _ []interface{}) error { return nil }
func (l *Log) RemoveEntity(_ *entity.Entity)                     {}

func (l *Log) DoTick(tick tick.Tick) (error, bool) {
	l.logger.LogTick(tick)

	l.sendLogAndSync()
	return nil, false
}

func (l *Log) sendLogAndSync() {
	l.logsIndex++
	if l.logsIndex >= l.sendLogsDelay {
		l.sendLog()
		l.logsIndex = 0
	} else if l.logsIndex == l.syncDelay {
		l.debugger.LogF("DoTick", "wait sync")
		<-l.ppWpLink.SyncCh
		l.debugger.LogF("DoTick", "sync get")
	}
}

func (l *Log) sendLog() {
	l.debugger.LogF("DoTick", "send logs")
	l.ppWpLink.LogsCh <- l.logger.GetLastBatch()
}
