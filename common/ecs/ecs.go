package ecs

import (
	"errors"

	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
)

var ErrNewEcsShouldBeWithAtLeastOneSystem = errors.New("should be at least one system")

type Ecs struct {
	systems    []System
	entityRepo EntityRepo
	debugger   nestedDebugger
}

func NewEcs(systems []System, entityRepo EntityRepo, debugger nestedDebugger) (*Ecs, error) {
	if len(systems) == 0 {
		return nil, ErrNewEcsShouldBeWithAtLeastOneSystem
	}
	return &Ecs{
		systems:    systems,
		entityRepo: entityRepo,
		debugger:   debugger,
	}, nil
}

func (ec *Ecs) Init(currentTick tick.Tick) {
	for _, s := range ec.systems {
		s.Init(currentTick)
	}
}

func (ec *Ecs) AddEntity(e entity.MaskedEntity) {
	ec.entityRepo.Add(e)
}

func (ec *Ecs) DoTick(currentTick tick.Tick) bool {
	var stop bool
	for _, s := range ec.systems {
		stopFromSystem := s.DoTick(currentTick)
		if stopFromSystem {
			stop = true
			ec.debugger.LogF("DoTick", "Stop from %s", s)
		}
	}
	return stop
}
