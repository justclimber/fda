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
}

type LogEntry struct {
	Tick tick.Tick
}

type Logs struct {
	Entries []LogEntry
	Batches []LogBatch
}

type TickComponent struct {
	Tick      tick.Tick
	Component component.Component
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
		Tick:      t,
		Component: c,
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
	logs *Logs
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
