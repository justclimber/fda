package servsystem

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
)

type TickLimiter struct {
	startTick tick.Tick
	limitTo   tick.Tick
}

func NewTickLimiter(startTick, limitTo tick.Tick) *TickLimiter {
	return &TickLimiter{
		startTick: startTick,
		limitTo:   limitTo,
	}
}

func (t *TickLimiter) String() string {
	return "TickLimiter"
}

func (t *TickLimiter) RequiredComponentKeys() []ecs.ComponentKey      { return nil }
func (t *TickLimiter) AddEntity(_ *ecs.Entity, _ []interface{}) error { return nil }
func (t *TickLimiter) RemoveEntity(_ *ecs.Entity)                     {}

func (t *TickLimiter) DoTick(tick tick.Tick) (error, bool) {
	return nil, tick-t.startTick >= t.limitTo-1
}
