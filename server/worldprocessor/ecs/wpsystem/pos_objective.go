package wpsystem

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type PosObjective struct {
	eId          entity.Id
	objectivePos fgeom.Point
	curPos       *wpcomponent.Position
}

func NewPosObjective(eId entity.Id, pos fgeom.Point) *PosObjective {
	return &PosObjective{
		eId:          eId,
		objectivePos: pos,
	}
}

func (p *PosObjective) String() string { return "PosObjective" }

func (p *PosObjective) Init() {}

func (p *PosObjective) RequiredComponentKeys() []component.Key {
	return []component.Key{wpcomponent.CPosition}
}

func (p *PosObjective) AddEntity(e *entity.Entity, in []interface{}) error {
	if e.Id != p.eId {
		return nil
	}
	c, ok := in[0].(*wpcomponent.Position)
	if !ok {
		return ErrInvalidComponent
	}
	p.curPos = c
	return nil
}

func (p *PosObjective) RemoveEntity(e *entity.Entity) {}

func (p *PosObjective) DoTick(_ tick.Tick) (error, bool) {
	return nil, p.curPos.Pos.EqualApprox(p.objectivePos, 0.1)
}
