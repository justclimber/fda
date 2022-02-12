package wpsystem

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/generated/wprepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type PosObjective struct {
	eId          entity.Id
	objectivePos fgeom.Point
	entityRepo   *wprepo.RepoForMask3
}

func NewPosObjective(repoForMask3 *wprepo.RepoForMask3, eId entity.Id, pos fgeom.Point) *PosObjective {
	p := &PosObjective{
		entityRepo:   repoForMask3,
		eId:          eId,
		objectivePos: pos,
	}
	p.entityRepo.InitRepoLink(p.mask())
	return p
}

func (p *PosObjective) String() string { return "PosObjective" }

func (p *PosObjective) Init(_ tick.Tick) {}

func (p *PosObjective) mask() component.Mask {
	return component.NewMask([]component.Key{wpcomponent.KeyMoving, wpcomponent.KeyPosition})
}

func (p *PosObjective) DoTick(_ tick.Tick) bool {
	stop := false
	p.entityRepo.Iterate(func(
		id entity.Id,
		pos wpcomponent.Position,
		mov wpcomponent.Moving,
	) (*wpcomponent.Position, *wpcomponent.Moving) {
		if id == p.eId && p.objectivePos.EqualApprox(pos.Pos, 0.1) {
			stop = true
		}
		return nil, nil
	})
	return stop
}
