package level

import (
	"fmt"

	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/server/ecs/servsystem"
	"github.com/justclimber/fda/server/levellog"
	"github.com/justclimber/fda/server/lpu"
)

type LpuAllocator struct {
	logger levellog.LevelLogger
}

func NewLpuAllocator(logger levellog.LevelLogger) *LpuAllocator {
	return &LpuAllocator{logger: logger}
}

func (a *LpuAllocator) GetLpuByEntity(_ *ecs.Entity) (*lpu.LevelProcessingUnit, error) {
	ec, err := ecs.NewEcs([]ecs.System{servsystem.NewMoving()})
	if err != nil {
		return nil, fmt.Errorf("faile to create ecs: %w", err)
	}
	return lpu.NewLpu(a.logger, ec), nil
}
