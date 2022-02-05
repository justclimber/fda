package worldprocessor

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/worldlog"
)

type WorldProcessor struct {
	logger        worldlog.WorldLogger
	ecs           *ecs.Ecs
	ppLink        *internalapi.PpWpLink
	logsIndex     int
	sendLogsDelay int
	syncDelay     int
	debugger      *debugger.Nested
}

func NewWorldProcessor(
	logger worldlog.WorldLogger,
	ecs *ecs.Ecs,
	ppLink *internalapi.PpWpLink,
	sendLogsDelay int,
	debugger *debugger.Nested,
) *WorldProcessor {
	return &WorldProcessor{
		logger:        logger,
		ecs:           ecs,
		ppLink:        ppLink,
		sendLogsDelay: sendLogsDelay,
		syncDelay:     sendLogsDelay - 2,
		debugger:      debugger,
	}
}

func (w *WorldProcessor) AddEntity(e *ecs.Entity) error {
	return w.ecs.AddEntity(e)
}

func (w *WorldProcessor) Run(currentTick tick.Tick) error {
	//w.logger.LogTick(currentTick)
	w.debugger.LogF("Run", "send logs on init")
	w.ppLink.LogsCh <- w.logger.GetLastBatch()
	for {
		w.debugger.LogF("Run", "[tick: %d]", currentTick)
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
	w.sendLogsAndSync()

	return nil, stop
}

func (w *WorldProcessor) sendLogsAndSync() {
	w.logsIndex++
	if w.logsIndex >= w.sendLogsDelay {
		w.debugger.LogF("Run", "send logs")
		w.ppLink.LogsCh <- w.logger.GetLastBatch()
		w.logsIndex = 0
	} else if w.logsIndex == w.syncDelay {
		w.debugger.LogF("Run", "wait sync")
		<-w.ppLink.SyncCh
		w.debugger.LogF("Run", "sync get")
	}
}
