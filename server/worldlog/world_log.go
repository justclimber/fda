package worldlog

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
)

type WorldLogger interface {
	LogTick(tick tick.Tick, eComps map[entity.Id]map[component.Key]component.Component)
	Logs() *Logs
	Count() int
}

type SingleTickComponent struct {
	Tick   tick.Tick
	EComps map[entity.Id]map[component.Key]component.Component
}

type Logs struct {
	Batches []Batch
}

type RepeatableComponent struct {
	TickFrom  tick.Tick
	TickTo    tick.Tick
	Component component.Component
}

type Batch struct {
	SingleTick []SingleTickComponent
	Repeatable map[entity.Id][]RepeatableComponent
}

func NewBatch() Batch {
	return Batch{Repeatable: map[entity.Id][]RepeatableComponent{}}
}

type Logger struct {
	logs                     *Logs
	curBatch                 Batch
	LastRepeatableComponents map[entity.Id]map[component.Key]int
}

func NewLogger() *Logger {
	return &Logger{
		logs:                     &Logs{Batches: []Batch{}},
		curBatch:                 NewBatch(),
		LastRepeatableComponents: map[entity.Id]map[component.Key]int{},
	}
}

func (l *Logger) LogTick(tick tick.Tick, eComps map[entity.Id]map[component.Key]component.Component) {
	l.curBatch.SingleTick = append(l.curBatch.SingleTick, SingleTickComponent{
		Tick:   tick,
		EComps: eComps,
	})
}

func (l *Logger) AddToCurBatch(t tick.Tick, id entity.Id, c component.Component) {
	tc := RepeatableComponent{
		TickFrom:  t,
		TickTo:    t,
		Component: c,
	}

	_, ok := l.curBatch.Repeatable[id]
	if !ok {
		l.curBatch.Repeatable[id] = []RepeatableComponent{tc}
		l.LastRepeatableComponents[id] = map[component.Key]int{
			c.Key(): 0,
		}
		return
	}

	lastComponents, ok := l.LastRepeatableComponents[id]
	if !ok {
		l.curBatch.Repeatable[id] = append(l.curBatch.Repeatable[id], tc)
		l.LastRepeatableComponents[id] = map[component.Key]int{
			c.Key(): len(l.curBatch.Repeatable[id]) - 1,
		}
		return
	}

	last, ok := lastComponents[c.Key()]
	if !ok || c != l.curBatch.Repeatable[id][last].Component {
		l.curBatch.Repeatable[id] = append(l.curBatch.Repeatable[id], tc)
		l.LastRepeatableComponents[id][c.Key()] = len(l.curBatch.Repeatable[id]) - 1
	} else {
		tc = l.curBatch.Repeatable[id][last]
		tc.TickTo = t
		l.curBatch.Repeatable[id][last] = tc
	}
}

func (l *Logger) RotateBatch() Batch {
	batch := l.curBatch
	l.logs.Batches = append(l.logs.Batches, batch)

	l.curBatch = NewBatch()
	l.LastRepeatableComponents = map[entity.Id]map[component.Key]int{}
	return batch
}

func (l *Logger) Logs() *Logs {
	return l.logs
}

// Count todo: remove in future
func (l *Logger) Count() int {
	return len(l.logs.Batches[0].SingleTick)
}
