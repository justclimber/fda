package wpsystem

import (
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/generated/wprepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type Moving struct {
	entityRepo *wprepo.RepoForMask3
}

func NewMoving(compIterator *wprepo.RepoForMask3) *Moving {
	return &Moving{entityRepo: compIterator}
}

func (m *Moving) String() string   { return "Moving" }
func (m *Moving) Init(_ tick.Tick) {}

func (m *Moving) DoTick(_ tick.Tick) bool {
	m.entityRepo.Iterate(func(
		_ entity.Id,
		p wpcomponent.Position,
		mov wpcomponent.Moving,
	) (*wpcomponent.Position, *wpcomponent.Moving) {
		if mov.D.Empty() {
			return nil, nil
		}
		return &wpcomponent.Position{Pos: p.Pos.Add(mov.D)}, nil
	})
	return false
}
