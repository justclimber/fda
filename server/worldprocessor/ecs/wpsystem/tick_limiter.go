package wpsystem

import (
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

func (t *TickLimiter) String() string   { return "TickLimiter" }
func (t *TickLimiter) Init(_ tick.Tick) {}

func (t *TickLimiter) DoTick(tick tick.Tick) bool {
	return tick-t.startTick >= t.limitTo-1
}
