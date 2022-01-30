package lpu

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/levellog"
)

type LevelProcessingUnit struct {
	logger levellog.LevelLogger
	ecs    *ecs.Ecs
}

func NewLpu(logger levellog.LevelLogger, ecs *ecs.Ecs) *LevelProcessingUnit {
	return &LevelProcessingUnit{
		logger: logger,
		ecs:    ecs,
	}
}

func (u *LevelProcessingUnit) AddEntity(e *ecs.Entity) error {
	return u.ecs.AddEntity(e)
}

func (u *LevelProcessingUnit) Run(currentTick tick.Tick) error {
	for {
		err, stop := u.doTick(currentTick)
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
		currentTick++
	}
}

func (u *LevelProcessingUnit) doTick(currentTick tick.Tick) (error, bool) {
	err, stop := u.ecs.DoTick(currentTick)
	if err != nil {
		return err, false
	}
	u.logger.LogTick(currentTick)

	return nil, stop
}

func (u *LevelProcessingUnit) Logger() levellog.LevelLogger {
	return u.logger
}
