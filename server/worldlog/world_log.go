package worldlog

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
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
	Batches []LogBatch
}

type TickComponent struct {
	Tick       tick.Tick
	Components map[component.Key]component.Component
}

type LogBatch struct {
	StartTick      tick.Tick
	EndTick        tick.Tick
	EntitiesLogs   map[entity.Id][]TickComponent
	LastComponents map[entity.Id]map[component.Key]component.Component
}

func NewLogBatch(t tick.Tick) LogBatch {
	return LogBatch{
		StartTick:      t,
		EntitiesLogs:   map[entity.Id][]TickComponent{},
		LastComponents: map[entity.Id]map[component.Key]component.Component{},
	}
}

func (l *LogBatch) Add(t tick.Tick, id entity.Id, c component.Component) {
	l.EndTick = t
	lastComponents, ok := l.LastComponents[id]
	tc := TickComponent{
		Tick: t,
		Components: map[component.Key]component.Component{
			c.Key(): c,
		},
	}
	if !ok {
		l.EntitiesLogs[id] = []TickComponent{tc}
		l.LastComponents[id] = map[component.Key]component.Component{
			wpcomponent.KeyMoving: c,
		}
		return
	}
	last, ok := lastComponents[wpcomponent.KeyMoving]
	if !ok || c != last {
		l.EntitiesLogs[id] = append(l.EntitiesLogs[id], tc)
		lastComponents[wpcomponent.KeyMoving] = c
	}
}

type Logger struct {
	logs              *Logs
	lastBatchLogIndex int
}

func NewLogger() *Logger {
	return &Logger{logs: &Logs{
		Entries: []LogEntry{},
		Batches: []LogBatch{},
	}}
}

func (l *Logger) LogTick(tick tick.Tick) {
	l.logs.Entries = append(l.logs.Entries, LogEntry{Tick: tick})
}

func (l *Logger) LogBatch(b LogBatch) {
	l.logs.Batches = append(l.logs.Batches, b)
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
