package wpsystem

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type PosObjective struct {
	eId          ecs.EntityId
	objectivePos fgeom.Point
	curPos       *wpcomponent.Position
}

func NewPosObjective(eId ecs.EntityId, pos fgeom.Point) *PosObjective {
	return &PosObjective{
		eId:          eId,
		objectivePos: pos,
	}
}

func (p *PosObjective) String() string {
	return "PosObjective"
}

func (p *PosObjective) RequiredComponentKeys() []ecs.ComponentKey {
	return []ecs.ComponentKey{wpcomponent.CPosition}
}

func (p *PosObjective) AddEntity(e *ecs.Entity, in []interface{}) error {
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

func (p *PosObjective) RemoveEntity(e *ecs.Entity) {}

func (p *PosObjective) DoTick(_ tick.Tick) (error, bool) {
	return nil, p.curPos.Pos.EqualApprox(p.objectivePos, 0.1)
}
