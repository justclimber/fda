package servsystem

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/ecs/servcomponent"
)

type components struct {
	Movable  *servcomponent.Movable
	Position *servcomponent.Position
}

type Moving struct {
	components map[ecs.EntityId]components
}

func NewMoving() *Moving {
	return &Moving{components: make(map[ecs.EntityId]components)}
}

func (m *Moving) AddEntity(e *ecs.Entity, in []interface{}) error {
	if len(in) != 2 {
		return ErrInvalidComponent
	}
	movable, ok1 := in[0].(*servcomponent.Movable)
	pos, ok2 := in[1].(*servcomponent.Position)
	if !ok1 || !ok2 {
		return ErrInvalidComponent
	}

	m.components[e.Id] = components{
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
		c.Position.Pos.X = c.Position.Pos.X + 1
	}
	return nil, false
}

func (m *Moving) RequiredComponentKeys() []ecs.ComponentKey {
	return []ecs.ComponentKey{
		servcomponent.CMovable,
		servcomponent.CPosition,
	}
}
