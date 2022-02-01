package worldlog

import (
	"github.com/justclimber/fda/common/tick"
)

type WorldLogger interface {
	LogTick(tick tick.Tick)
	Logs() []*LogEntry
}

type LogEntry struct {
	Tick tick.Tick
}

type WorldLog struct {
	logs []*LogEntry
}

func NewWorldLog() *WorldLog {
	return &WorldLog{}
}

func (l *WorldLog) LogTick(tick tick.Tick) {
	l.logs = append(l.logs, &LogEntry{Tick: tick})
}

func (l *WorldLog) Logs() []*LogEntry {
	return l.logs
}
