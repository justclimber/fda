package worldlog

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
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
	Batches []LogBatch
}

type TickComponent struct {
	TickFrom  tick.Tick
	TickTo    tick.Tick
	Component component.Component
}

type LogBatch struct {
	EntitiesLogs map[entity.Id][]TickComponent
}

func NewLogBatch() LogBatch {
	return LogBatch{EntitiesLogs: map[entity.Id][]TickComponent{}}
}

type Logger struct {
	logs           *Logs
	curLogBatch    LogBatch
	LastComponents map[entity.Id]map[component.Key]int
}

func NewLogger() *Logger {
	return &Logger{
		logs: &Logs{
			Entries: []LogEntry{},
			Batches: []LogBatch{},
		},
		curLogBatch:    NewLogBatch(),
		LastComponents: map[entity.Id]map[component.Key]int{},
	}
}

func (l *Logger) LogTick(tick tick.Tick) {
	l.logs.Entries = append(l.logs.Entries, LogEntry{Tick: tick})
}

func (l *Logger) AddToCurLogBatch(t tick.Tick, id entity.Id, c component.Component) {
	tc := TickComponent{
		TickFrom:  t,
		TickTo:    t,
		Component: c,
	}

	_, ok := l.curLogBatch.EntitiesLogs[id]
	if !ok {
		l.curLogBatch.EntitiesLogs[id] = []TickComponent{tc}
		l.LastComponents[id] = map[component.Key]int{
			c.Key(): 0,
		}
		return
	}

	lastComponents, ok := l.LastComponents[id]
	if !ok {
		l.curLogBatch.EntitiesLogs[id] = append(l.curLogBatch.EntitiesLogs[id], tc)
		l.LastComponents[id] = map[component.Key]int{
			c.Key(): len(l.curLogBatch.EntitiesLogs[id]) - 1,
		}
		return
	}

	last, ok := lastComponents[c.Key()]
	if !ok || c != l.curLogBatch.EntitiesLogs[id][last].Component {
		l.curLogBatch.EntitiesLogs[id] = append(l.curLogBatch.EntitiesLogs[id], tc)
		l.LastComponents[id][c.Key()] = len(l.curLogBatch.EntitiesLogs[id]) - 1
	} else {
		tc = l.curLogBatch.EntitiesLogs[id][last]
		tc.TickTo = t
		l.curLogBatch.EntitiesLogs[id][last] = tc
	}
}

func (l *Logger) RotateLogBatch() LogBatch {
	batch := l.curLogBatch
	l.logs.Batches = append(l.logs.Batches, batch)

	l.curLogBatch = NewLogBatch()
	return batch
}

func (l *Logger) Logs() *Logs {
	return l.logs
}

func (l *Logger) Count() int {
	return len(l.logs.Entries)
}
