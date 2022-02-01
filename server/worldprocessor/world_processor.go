package worldprocessor

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldlog"
)

type WorldProcessor struct {
	logger worldlog.WorldLogger
	ecs    *ecs.Ecs
}

func NewWorldProcessor(logger worldlog.WorldLogger, ecs *ecs.Ecs) *WorldProcessor {
	return &WorldProcessor{
		logger: logger,
		ecs:    ecs,
	}
}

func (w *WorldProcessor) AddEntity(e *ecs.Entity) error {
	return w.ecs.AddEntity(e)
}

func (w *WorldProcessor) Run(currentTick tick.Tick) error {
	for {
		err, stop := w.doTick(currentTick)
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
		currentTick++
	}
}

func (w *WorldProcessor) doTick(currentTick tick.Tick) (error, bool) {
	err, stop := w.ecs.DoTick(currentTick)
	if err != nil {
		return err, false
	}
	w.logger.LogTick(currentTick)

	return nil, stop
}
