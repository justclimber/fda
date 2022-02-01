package world

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/server/ecs/servcomponent"
)

type World struct{}

func NewWorld() *World {
	return &World{}
}

func NewPlayerEntity(id ecs.EntityId, p *servcomponent.Player) *ecs.Entity {
	return &ecs.Entity{
		Id: id,
		Components: map[ecs.ComponentKey]interface{}{
			servcomponent.CPlayer: p,
		},
	}
}

func (w *World) RegisterNewEntity(*ecs.Entity) error {
	return nil
}

func (w *World) AllocateEntity(*ecs.Entity) error {
	return nil
}
