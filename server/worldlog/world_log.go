package worldlog

import (
	"github.com/justclimber/fda/common/tick"
)

type WorldLogger interface {
	LogTick(tick tick.Tick)
	Logs() *Logs
	Count() int
	GetLastBatch() *Logs
}

type LogEntry struct {
	Tick tick.Tick
}

type Logs struct {
	Entries []LogEntry
}

type Logger struct {
	logs              *Logs
	lastBatchLogIndex int
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

func (l *Logger) GetLastBatch() *Logs {
	if len(l.logs.Entries) == 0 {
		return nil
	}
	i := l.lastBatchLogIndex
	l.lastBatchLogIndex = len(l.logs.Entries) - 1
	return &Logs{Entries: l.logs.Entries[i:]}
}
