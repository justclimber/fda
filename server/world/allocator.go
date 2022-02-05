package world

import (
	"github.com/justclimber/fda/server/worldlog"
)

type Allocator struct {
	logger worldlog.WorldLogger
}

func NewAllocator(logger worldlog.WorldLogger) *Allocator {
	return &Allocator{logger: logger}
}

//func (a *Allocator) GetLpuByEntity(_ *ecs.Entity) (*worldprocessor.WorldProcessor, error) {
//	ec, err := ecs.NewEcs([]ecs.System{wpsystem.NewMoving()})
//	if err != nil {
//		return nil, fmt.Errorf("faile to create ecs: %w", err)
//	}
//	return worldprocessor.NewWorldProcessor(a.logger, ec, nil), nil
//}
