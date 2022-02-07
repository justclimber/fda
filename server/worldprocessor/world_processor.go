package worldprocessor

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
)

type WorldProcessor struct {
	ecs      *ecs.Ecs
	ppLink   *internalapi.PpWpLink
	debugger *debugger.Nested
}

func NewWorldProcessor(
	ecs *ecs.Ecs,
	ppLink *internalapi.PpWpLink,
	debugger *debugger.Nested,
) *WorldProcessor {
	return &WorldProcessor{
		ecs:      ecs,
		ppLink:   ppLink,
		debugger: debugger,
	}
}

func (w *WorldProcessor) AddEntity(e *entity.Entity) error {
	return w.ecs.AddEntity(e)
}

func (w *WorldProcessor) Run(currentTick tick.Tick) error {
	w.ecs.Init(currentTick)
	for {
		w.debugger.LogF("Run", "[tick: %d]", currentTick)
		err, stop := w.ecs.DoTick(currentTick)
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
