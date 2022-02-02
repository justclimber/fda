package worldprocessor

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/worldlog"
)

type WorldProcessor struct {
	logger worldlog.WorldLogger
	ecs    *ecs.Ecs
	ppLink *internalapi.PpWpLink
}

func NewWorldProcessor(logger worldlog.WorldLogger, ecs *ecs.Ecs, ppLink *internalapi.PpWpLink) *WorldProcessor {
	return &WorldProcessor{
		logger: logger,
		ecs:    ecs,
		ppLink: ppLink,
	}
}

func (w *WorldProcessor) AddEntity(e *ecs.Entity) error {
	return w.ecs.AddEntity(e)
}

func (w *WorldProcessor) Run(currentTick tick.Tick) error {
	w.ppLink.LogsCh <- w.logger.GetLastBatch()
	for {
		err, stop := w.doTick(currentTick)
		if err != nil {
			return err
		}
		if stop {
			w.ppLink.DoneCh <- true
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

	w.ppLink.LogsCh <- w.logger.GetLastBatch()

	return nil, stop
}
