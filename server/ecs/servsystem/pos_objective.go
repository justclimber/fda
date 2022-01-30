package servsystem

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/ecs/servcomponent"
)

type PosObjective struct {
	eId          ecs.EntityId
	objectivePos fgeom.Point
	curPos       *servcomponent.Position
}

func NewPosObjective(eId ecs.EntityId, pos fgeom.Point) *PosObjective {
	return &PosObjective{
		eId:          eId,
		objectivePos: pos,
	}
}

func (p *PosObjective) AddEntity(e *ecs.Entity, in []interface{}) error {
	if e.Id != p.eId {
		return nil
	}
	c, ok := in[0].(*servcomponent.Position)
	if !ok {
		return ErrInvalidComponent
	}
	p.curPos = c
	return nil
}

func (p *PosObjective) RemoveEntity(e *ecs.Entity) {
}

func (p *PosObjective) DoTick(_ tick.Tick) (error, bool) {
	return nil, p.curPos.Pos.X == p.objectivePos.X
}

func (p *PosObjective) RequiredComponentKeys() []ecs.ComponentKey {
	return []ecs.ComponentKey{servcomponent.CPosition}
}
