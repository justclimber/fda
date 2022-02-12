package wpsystem

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wprepo"
)

type Moving struct {
	entityRepo *wprepo.RepoForMask3
}

func NewMoving(compIterator *wprepo.RepoForMask3) *Moving {
	m := &Moving{entityRepo: compIterator}
	m.entityRepo.InitRepoLink(m.mask())
	return m
}

func (m *Moving) String() string { return "Moving" }

func (m *Moving) Init(_ tick.Tick) {}

func (m *Moving) mask() component.Mask {
	return component.NewMask([]component.Key{wpcomponent.KeyMoving, wpcomponent.KeyPosition})
}

func (m *Moving) DoTick(_ tick.Tick) bool {
	m.entityRepo.Iterate(func(
		_ entity.Id,
		mov wpcomponent.Moving,
		p wpcomponent.Position,
	) (*wpcomponent.Moving, *wpcomponent.Position) {
		if mov.D.Empty() {
			return nil, nil
		}
		return nil, &wpcomponent.Position{Pos: p.Pos.Add(mov.D)}
	})
	return false
}
