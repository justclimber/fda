package wpsystem

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type movingCs struct {
	Movable  wpcomponent.Movable
	Position *wpcomponent.Position
}

type Moving struct {
	components map[ecs.EntityId]movingCs
}

func NewMoving() *Moving {
	return &Moving{components: make(map[ecs.EntityId]movingCs)}
}

func (m *Moving) String() string {
	return "Moving"
}

func (m *Moving) RequiredComponentKeys() []ecs.ComponentKey {
	return []ecs.ComponentKey{
		wpcomponent.CMovable,
		wpcomponent.CPosition,
	}
}

func (m *Moving) AddEntity(e *ecs.Entity, in []interface{}) error {
	if len(in) != 2 {
		return ErrInvalidComponent
	}
	movable, ok1 := in[0].(wpcomponent.Movable)
	pos, ok2 := in[1].(*wpcomponent.Position)
	if !ok1 || !ok2 {
		return ErrInvalidComponent
	}

	m.components[e.Id] = movingCs{
		Movable:  movable,
		Position: pos,
	}
	return nil
}

func (m *Moving) RemoveEntity(e *ecs.Entity) {
	delete(m.components, e.Id)
}

func (m *Moving) DoTick(_ tick.Tick) (error, bool) {
	for _, c := range m.components {
		c.Movable.Move(c.Position.Pos)
	}
	return nil, false
}
