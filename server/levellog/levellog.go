package levellog

import (
	"github.com/justclimber/fda/common/tick"
)

type LevelLogger interface {
	LogTick(tick tick.Tick)
	Logs() []*LogEntry
}

type LogEntry struct {
	Tick tick.Tick
}

type LevelLog struct {
	logs []*LogEntry
}

func NewLevelLog() *LevelLog {
	return &LevelLog{}
}

func (l *LevelLog) LogTick(tick tick.Tick) {
	l.logs = append(l.logs, &LogEntry{Tick: tick})
}

func (l *LevelLog) Logs() []*LogEntry {
	return l.logs
}
