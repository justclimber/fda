package worldlog

import (
	"github.com/justclimber/fda/common/tick"
)

type WorldLogger interface {
	LogTick(tick tick.Tick)
	Logs() *Logs
	Count() int
}

type LogEntry struct {
	Tick tick.Tick
}

type Logs struct {
	Entries []LogEntry
}

type Logger struct {
	logs *Logs
}

func NewLogger() *Logger {
	return &Logger{logs: &Logs{
		Entries: []LogEntry{},
	}}
}

func (l *Logger) LogTick(tick tick.Tick) {
	l.logs.Entries = append(l.logs.Entries, LogEntry{Tick: tick})
}

func (l *Logger) Logs() *Logs {
	return l.logs
}

func (l *Logger) Count() int {
	return len(l.logs.Entries)
}
