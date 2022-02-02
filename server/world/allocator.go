package world

import (
	"fmt"

	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/server/ecs/servsystem"
	"github.com/justclimber/fda/server/worldlog"
	"github.com/justclimber/fda/server/worldprocessor"
)

type Allocator struct {
	logger worldlog.WorldLogger
}

func NewAllocator(logger worldlog.WorldLogger) *Allocator {
	return &Allocator{logger: logger}
}

func (a *Allocator) GetLpuByEntity(_ *ecs.Entity) (*worldprocessor.WorldProcessor, error) {
	ec, err := ecs.NewEcs([]ecs.System{servsystem.NewMoving()})
	if err != nil {
		return nil, fmt.Errorf("faile to create ecs: %w", err)
	}
	return worldprocessor.NewWorldProcessor(a.logger, ec, nil), nil
}
